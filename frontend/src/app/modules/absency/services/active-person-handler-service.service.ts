import { DestroyRef, Injectable, effect, inject, signal } from '@angular/core';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { catchError, filter, forkJoin, map, of, switchMap, throwError } from 'rxjs';
import { AbsenceForPerson, AbsenceReponse } from '../interfaces/absence';
import { groupDatesToRanges } from '../functions/group-dates-to-ranges.function';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { NotificationService } from '@app/core/services/notification.service';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import CalendarDayEventObject from 'js-year-calendar/dist/interfaces/CalendarDayEventObject';
import CalendarDataSourceElement from 'js-year-calendar/dist/interfaces/CalendarDataSourceElement';
import { messages } from '@app/constants/messages';
import { CreateAbsencyDialogComponent } from '../components/create-absency-dialog/create-absency-dialog.component';
import { formatDateToDateString } from '../../../shared/functions/format-date-to-string.function';
import { ConfirmationDialogComponent, ConfirmationDialogComponentData } from '@app/shared/components/confirmation-dialog/confirmation-dialog.component';
@Injectable({
  providedIn: null,
})
export class ActivePersonHandlerServiceService {
  private destroyRef$ = inject(DestroyRef);
  private notificationService = inject(NotificationService);
  private _personAPIService = inject(PersonAPIService);
  dialog = inject(MatDialog);

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

  /* 
  The following code is responsible for keeping track of the absences for a given person and year 
  */
  private _absences = signal<AbsenceForPerson[]>([]);
  set absences(absences: AbsenceReponse[]) {
    // group by reason into a map
    const absencesByReason = absences.reduce((acc, absence) => {
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
    }, new Map<string, { date: Date; created_at: Date }[]>());

    // group each reason into ranges of absences instead of individual per date absences
    const absencesResult: AbsenceForPerson[] = [];
    absencesByReason.forEach((absencesArray, reason) => {
      absencesArray.sort((a, b) => a.date.getTime() - b.date.getTime());
      absencesResult.push(...groupDatesToRanges(absencesArray, reason));
    });

    this._absences.set(absencesResult);
  }
  get absences$() {
    return this._absences();
  }

  constructor() {
    effect(
      () => {
        of(this.activeYear$)
          .pipe(
            takeUntilDestroyed(this.destroyRef$),
            filter(() => !!this.activePerson$),
            map((year) => ({
              startDate: year + '-01-01',
              endDate: year + '-12-31',
              personId: this.activePerson$!.id,
            })),
            switchMap(({ startDate, endDate, personId }) => this._personAPIService.getAbsencyForPersonInRange(personId, startDate, endDate).pipe(catchError((err) => throwError(() => err)))),
            map((resp) => resp.data),
            map((absences) => (absences ? absences : [])),
          )
          .subscribe((absences) => {
            this.absences = absences;
          });
      },
      { allowSignalWrites: true },
    );
  }

  private _addAbsency(e: CalendarDayEventObject<CalendarDataSourceElement>) {
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
          // prevent race condition if the start and end date are the same
          if (result.endDate === e.date) return { dates: [e.date], reason: result.reason };

          // Calculate all dates between the start and end date
          for (let date = result.endDate; date >= e.date; date.setDate(date.getDate() - 1)) {
            dates.push(new Date(date));
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
        switchMap(() =>
          this._personAPIService.getAbsencyForPersonInRange(this.activePerson$!.id, this.activeYear$ + '-01-01', this.activeYear$ + '-12-31').pipe(catchError((err) => throwError(() => err))),
        ),
        map((resp) => resp.data),
        map((data) => (data ? data : [])),
      )
      .subscribe({
        next: (absences) => {
          this.notificationService.infoMessage(messages.ABSENCY.CREATED);
          this.absences = absences;
        },
        error: () => {
          this.notificationService.warnMessage(messages.GENERAL.HTTP_ERROR.SERVER_ERROR);
        },
      });
  }

  private _removeAbsency(e: CalendarDayEventObject<CalendarDataSourceElement>) {
    const dialogData: ConfirmationDialogComponentData = {
      title: this.activePerson$?.id ?? '',
      confirmationMessage: messages.ABSENCY.DELETE_CONFIRMATION,
    };

    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = dialogData;
    dialogConfig.enterAnimationDuration = 300;
    dialogConfig.exitAnimationDuration = 300;

    const dialogRef = this.dialog.open(ConfirmationDialogComponent, dialogConfig);
    dialogRef
      .afterClosed()
      .pipe(
        // We dont need to do anything if the dialog was closed without a result
        filter((result) => result !== null && result !== undefined && result === true),
        map(() => ({ personID: this.activePerson$?.id ?? '', date: e.date })),
        map(({ personID, date }) => ({ personID, date: formatDateToDateString(date) })),
        switchMap(({ personID, date }) => this._personAPIService.removeAbsencyFromPerson(personID, date).pipe(catchError((err) => throwError(() => err)))),
        switchMap(() =>
          this._personAPIService.getAbsencyForPersonInRange(this.activePerson$!.id, this.activeYear$ + '-01-01', this.activeYear$ + '-12-31').pipe(catchError((err) => throwError(() => err))),
        ),
        map((resp) => resp.data),
        map((data) => (data ? data : [])),
      )
      .subscribe({
        next: (absences) => {
          this.notificationService.infoMessage(messages.ABSENCY.DELETED);
          this.absences = absences;
        },
        error: () => {
          this.notificationService.warnMessage(messages.GENERAL.HTTP_ERROR.SERVER_ERROR);
        },
      });
  }

  handleDayClick(e: CalendarDayEventObject<CalendarDataSourceElement>) {
    /** Open dialog and display the interface to create an absency */
    if (e.events.length > 0) {
      this._removeAbsency(e);
      return;
    }
    this._addAbsency(e);
  }

  // functions for unit testing
  removeAbsency(e: CalendarDayEventObject<CalendarDataSourceElement>) {
    this._removeAbsency(e);
  }
  addAbsency(e: CalendarDayEventObject<CalendarDataSourceElement>) {
    this._addAbsency(e);
  }
}
