import { Injectable, inject } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { NotificationService } from '@app/core/services/notification.service';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { Subject, catchError, filter, forkJoin, from, map, mergeMap, of, reduce, switchMap, tap, throwError } from 'rxjs';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
const getWeekFromDate = (date: Date) => {
  // Create a new date object from the input date
  const currentDate = new Date(date.getTime());

  // Calculate the nearest Monday
  while (currentDate.getDay() !== 1) {
    currentDate.setDate(currentDate.getDate() - 1);
  }

  // Create an array to store all dates of the week
  const week = [new Date(currentDate.getTime())];

  // Add the next 6 days to the array
  for (let i = 1; i <= 6; i++) {
    currentDate.setDate(currentDate.getDate() + 1);
    week.push(new Date(currentDate.getTime()));
  }

  return week;
};

type ActiveWeek = {
  department: string;
  dates: Date[];
};

@Injectable({
  providedIn: null,
})
export class PlannerStateHandlerService {
  private notificationService = inject(NotificationService);
  private personAPIService = inject(PersonAPIService);
  private timetableDataContainerService = inject(TimetableDataContainerService);
  private personDataContainerService = inject(PersonDataContainerService);
  private workdayAPIService = inject(WorkdayAPIService);

  // This keeps track of the active department currently being viewed
  setActiveView(department: string, date: Date) {
    this._activeViewTrackerSubject.next({ department, date });
  }
  private _activeViewTrackerSubject = new Subject<{
    department: string;
    date: Date;
  }>();

  constructor() {
    // this keeps track of the active week currently being viewed
    // this is a computed property which fetches the workdays upon receiving activeWeek signal
    this._activeViewTrackerSubject
      .pipe(
        takeUntilDestroyed(),
        map(({ department, date }) => ({
          department,
          dates: getWeekFromDate(date),
        })),
        tap((activeWeek) => {
          this.timetableDataContainerService.weekdays = activeWeek.dates;
        }),
        // filter out null values
        filter((activeWeek): activeWeek is ActiveWeek => !!activeWeek),
        // date to YYYY-MM-DD format
        map(({ department, dates }) => ({
          department,
          dates: dates.map((date) => date.toISOString().split('T')[0]),
        })),
        // call API for each date in dates and return the workdays for each date
        mergeMap(({ department, dates }) =>
          // convert the dates array to an observable
          from(dates).pipe(
            // fetch workdays for each date
            mergeMap((date) =>
              this.workdayAPIService.getWorkdays(department, date).pipe(
                map((resp) => resp.data), // map the response to the data property
                catchError(() => of([])), // if there is an error, return an empty array
              ),
            ),
            reduce((acc, workdays) => [...acc, ...workdays], [] as WorkdayTimeslot[]), // reduce the workdays into a single array
          ),
        ),
        map((workdays) => workdays.sort((a, b) => a.workplace.id.localeCompare(b.workplace.id))),
      )
      .subscribe((workdays) => {
        // update the workdays in the service
        this.timetableDataContainerService.workdays = workdays;
      });
    this._activeViewTrackerSubject
      .pipe(
        takeUntilDestroyed(),
        map(({ department }) => department),
        switchMap((department) => this.personAPIService.getPersons(department).pipe(catchError((err) => throwError(() => err)))),
        map((response) => response.data),
      )
      .subscribe({
        next: (persons) => {
          this.personDataContainerService.persons = persons;
        },
      });
  }

  assignPersonToTimelots(event: { person: PersonWithMetadata; timeslots: WorkdayTimeslot[]; actionToBeExecutedOnFailedValidation?: () => void }) {
    /**
     * This function is responsible for validating the timeslots and updating the timeslots in the service. Furthermore it uses the API to update the timeslots on the server.
     * Validation requirements:
     *  - is qualified for the workplace
     *  - is not absent on the day
     *  - is not already assigned to a timeslot on the same day
     *
     */
    const { person, timeslots, actionToBeExecutedOnFailedValidation } = event;
    forkJoin(timeslots.map((timeslot) => this._assignPersonToTimeslot(person, timeslot, actionToBeExecutedOnFailedValidation))).subscribe((results) => {
      if (results.every((result) => result)) this.notificationService.infoMessage('Person erfolgreich zugeordnet');
    });
  }

  private _assignPersonToTimeslot(person: PersonWithMetadata, workdayTimeslot: WorkdayTimeslot, actionToBeExecutedOnFailedValidation?: () => void) {
    return of(workdayTimeslot).pipe(
      // check if the person is qualified for the workplace
      map((ts) => (person.workplaces ?? []).map((wp) => wp.id).includes(ts.workplace.id)),
      tap((isQualified) => {
        if (!isQualified) throw new Error('Person ist nicht für diesen Arbeitsplatz qualifiziert');
      }),
      // check if the person is absent on the day
      switchMap(() => this.personAPIService.isAbsentOnDate(person.id, workdayTimeslot.date).pipe(catchError(() => throwError(() => new Error('Datenbankfehler'))))),
      tap((isAbsent) => {
        if (isAbsent) throw new Error('Person ist an diesem Tag abwesend (Krank, Urlaub, etc.)');
      }),
      // check if the person is already assigned to a timeslot on the same day
      switchMap(() => this.workdayAPIService.getWorkdays(workdayTimeslot.department.id, workdayTimeslot.date).pipe(catchError(() => throwError(() => new Error('Datenbankfehler'))))),
      map((resp) => resp.data),
      map((workdays) =>
        workdays
          .filter((wd) => wd.date === workdayTimeslot.date)
          .map((wd) => wd.person?.id)
          .filter((id) => !!id)
          .includes(person.id),
      ),
      tap((isAssigned) => {
        if (isAssigned) throw new Error('Person ist bereits einem anderen Timeslot zugeordnet');
      }),
      // check if person is present on the day
      map(() => {
        return (person.weekdays ?? []).map((wd) => wd.id).includes(workdayTimeslot.weekday);
      }),
      tap((isPresent) => {
        if (!isPresent) throw new Error('Person ist an diesem Tag nicht anwesend');
      }),
      map(() => true),
      catchError((err: Error) => {
        this.notificationService.warnMessage(err.message);
        actionToBeExecutedOnFailedValidation?.();
        return of(false);
      }),
      filter((isValid) => isValid),
      // update the timeslot in the service
      switchMap(() => {
        return this.workdayAPIService.assignPerson(workdayTimeslot.department.id, workdayTimeslot.date, workdayTimeslot.workplace.id, workdayTimeslot.timeslot.name, person.id).pipe(
          catchError((err) => throwError(() => err)),
          map((resp) => resp.data),
        );
      }),
      map(() => {
        workdayTimeslot.person = person;
        return of(true);
      }),
      catchError((err) => {
        console.error(err);
        return of(false);
      }),
    );
  }
}
