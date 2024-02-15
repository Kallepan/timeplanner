import { DestroyRef, Injectable, effect, inject, signal } from '@angular/core';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { catchError, filter, forkJoin, map, of, switchMap, throwError } from 'rxjs';
import { Absence } from '../interfaces/absence';
import { groupDatesToRanges } from '../functions/group-dates-to-ranges.function';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { NotificationService } from '@app/core/services/notification.service';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import CalendarDayEventObject from 'js-year-calendar/dist/interfaces/CalendarDayEventObject';
import CalendarDataSourceElement from 'js-year-calendar/dist/interfaces/CalendarDataSourceElement';
import { messages } from '@app/constants/messages';
import { CreateAbsencyDialogComponent } from '../components/create-absency-dialog/create-absency-dialog.component';
import { formatDateToDateString } from '../../../shared/functions/format-date-to-string.function';
@Injectable({
  providedIn: null,
})
export class ActivePersonHandlerServiceService {
  private destroyRef$ = inject(DestroyRef);
  private notificationService = inject(NotificationService);
  private _personAPIService = inject(PersonAPIService);
  private dialog = inject(MatDialog);

  private _activePerson = signal<PersonWithMetadata | null>(null);
  set activePerson(person: PersonWithMetadata) {
    this._activePerson.set(person);
  }
  get activePerson$() {
    return this._activePerson();
  }

  private _activeYear = signal<number>(new Date().getFullYear());
  set activeYear(year: number) {
    this._activeYear.set(year);
  }
  get activeYear$() {
    return this._activeYear();
  }

  private _absences = signal<Absence[]>([]);
  get absences$() {
    return this._absences();
  }

  constructor() {
    effect(
      () => {
        of(this.activeYear$)
          .pipe(
            takeUntilDestroyed(this.destroyRef$),
            map((year) => ({
              startDate: new Date(year, 0, 1),
              endDate: new Date(year, 11, 31),
            })),
            filter(() => !!this.activePerson$),
            map(({ startDate, endDate }) => ({
              startDate: formatDateToDateString(startDate),
              endDate: formatDateToDateString(endDate),
              personId: this.activePerson$!.id,
            })),
            switchMap(({ startDate, endDate, personId }) => this._personAPIService.getAbsencyForPersonInRange(personId, startDate, endDate).pipe(catchError((err) => throwError(() => err)))),
            map((resp) => resp.data),
            filter((absences) => !!absences),
            // group by reason into a map
            map((absences) =>
              absences.reduce((acc, absence) => {
                const groupedByArray = acc.get(absence.reason);
                if (!groupedByArray) {
                  acc.set(absence.reason, [
                    {
                      date: new Date(absence.date),
                      created_at: absence.created_at,
                    },
                  ]);
                  return acc;
                }

                groupedByArray.push({
                  date: new Date(absence.date),
                  created_at: absence.created_at,
                });
                return acc;
              }, new Map<string, { date: Date; created_at: Date }[]>()),
            ),
            // group each reason into ranges of absences instead of individual per date absences
            map((absencesMap) => {
              const absences: Absence[] = [];

              absencesMap.forEach((absencesArray, reason) => {
                absencesArray.sort((a, b) => a.date.getTime() - b.date.getTime());
                absences.push(...groupDatesToRanges(absencesArray, reason));
              });

              return absences;
            }),
          )
          .subscribe((absences) => {
            this._absences.set(absences);
          });
      },
      { allowSignalWrites: true },
    );
  }

  handleDayClick(e: CalendarDayEventObject<CalendarDataSourceElement>) {
    /** Open dialog and display the interface to create an absency */
    if (e.events.length > 0) {
      this.notificationService.infoMessage(messages.ABSENCY.ALREADY_EXISTS);
      return;
    }

    const dialogData = {
      personID: this.activePerson$?.id ?? '',
      startDate: e.date,
    };

    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = dialogData;
    dialogConfig.enterAnimationDuration = 300;
    dialogConfig.exitAnimationDuration = 300;

    const dialogRef = this.dialog.open(CreateAbsencyDialogComponent, dialogConfig);
    dialogRef
      .afterClosed()
      .pipe(
        filter((result) => result !== null && result !== undefined),
        map((result) => result as { endDate: Date; reason: string }),
        // calculate dates between start and end date
        map((result) => {
          const dates: Date[] = [];
          // We need to clone the date object to avoid modifying the original date
          const currentDate = new Date(result.endDate);
          const endDate = new Date(result.endDate);
          while (currentDate <= endDate) {
            dates.push(new Date(currentDate));
            currentDate.setDate(currentDate.getDate() + 1);
          }
          return { dates, reason: result.reason };
        }),
        map(({ dates, reason }) => {
          const formattedDates = dates.map((date) => formatDateToDateString(date));

          return { dates: formattedDates, reason };
        }),
        // call the api for each date
        switchMap(({ dates, reason }) => {
          const obs = dates.map((date) => this._personAPIService.addAbsencyToPerson(this.activePerson$!.id, date, reason).pipe(catchError((err) => throwError(() => err))));
          return forkJoin(obs);
        }),
      )
      .subscribe({
        next: () => {
          this.notificationService.infoMessage(messages.ABSENCY.CREATED);
        },
        error: () => {
          this.notificationService.warnMessage(messages.GENERAL.HTTP_ERROR.SERVER_ERROR);
        },
      });
  }
}
