import { Injectable, inject } from '@angular/core';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { Subject, catchError, filter, from, map, mergeMap, of, reduce, tap } from 'rxjs';
import { TimetableDataContainerService } from '../../../shared/services/timetable-data-container.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

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
export class ViewerStateHandlerService {
  // inject the services here
  private timetableDataContainerService = inject(TimetableDataContainerService);
  private workdayAPIService = inject(WorkdayAPIService);

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
  }
}
