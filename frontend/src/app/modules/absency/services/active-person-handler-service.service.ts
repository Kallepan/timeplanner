import { DestroyRef, Injectable, effect, inject, signal } from '@angular/core';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { catchError, filter, map, of, switchMap, throwError } from 'rxjs';
import { Absence } from '../interfaces/absence';
import { groupDatesToRanges } from '../functions/group-dates-to-ranges.function';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
@Injectable({
  providedIn: null,
})
export class ActivePersonHandlerServiceService {
  private destroyRef$ = inject(DestroyRef);
  private _personAPIService = inject(PersonAPIService);

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
              startDate: startDate.toISOString().split('T')[0],
              endDate: endDate.toISOString().split('T')[0],
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
}
