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
import { messages } from '@app/constants/messages';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { EditTextareaDialogComponent, EditTextareaDialogData } from '@app/shared/components/edit-textarea-dialog/edit-textarea-dialog.component';
import { FormControl } from '@angular/forms';
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
  // dialog
  private dialog = inject(MatDialog);

  // services
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

  assignPersonToTimelots(person: PersonWithMetadata, timeslots: WorkdayTimeslot[], actionToBeExecutedOnFailedValidation?: () => void) {
    /**
     * This function is responsible for validating the timeslots and updating the timeslots in the service. Furthermore it uses the API to update the timeslots on the server.
     * Validation requirements:
     *  - is qualified for the workplace
     *  - is not absent on the day
     *  - is not already assigned to a timeslot on the same day
     *
     */
    forkJoin(timeslots.map((timeslot) => this._assignPersonToTimeslot(person, timeslot, actionToBeExecutedOnFailedValidation))).subscribe((results) => {
      if (results.every((result) => result)) this.notificationService.infoMessage(messages.PLANNER.TIMESLOT_ASSIGNMENT.SUCCESS);
    });
  }

  private _assignPersonToTimeslot(person: PersonWithMetadata, workdayTimeslot: WorkdayTimeslot, actionToBeExecutedOnFailedValidation?: () => void) {
    return of(workdayTimeslot).pipe(
      // check if the person is qualified for the workplace
      map((ts) => (person.workplaces ?? []).map((wp) => wp.id).includes(ts.workplace.id)),
      tap((isQualified) => {
        if (!isQualified) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_NOT_QUALIFIED);
      }),
      // check if the person is absent on the day
      switchMap(() => this.personAPIService.isAbsentOnDate(person.id, workdayTimeslot.date).pipe(catchError(() => throwError(() => new Error(messages.GENERAL.HTTP_ERROR.SERVER_ERROR))))),
      tap((isAbsent) => {
        if (isAbsent) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_ABSENT);
      }),
      // check if the person is already assigned to a timeslot on the same day
      switchMap(() =>
        this.workdayAPIService.getWorkdays(workdayTimeslot.department.id, workdayTimeslot.date).pipe(catchError(() => throwError(() => new Error(messages.GENERAL.HTTP_ERROR.SERVER_ERROR)))),
      ),
      map((resp) => resp.data),
      map((workdays) =>
        workdays
          .filter((wd) => wd.date === workdayTimeslot.date)
          .map((wd) => wd.person?.id)
          .filter((id) => !!id)
          .includes(person.id),
      ),
      tap((isAssigned) => {
        if (isAssigned) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_ALREADY_ASSIGNED);
      }),
      // check if person is present on the day
      map(() => {
        return (person.weekdays ?? []).map((wd) => wd.id).includes(workdayTimeslot.weekday);
      }),
      tap((isPresent) => {
        if (!isPresent) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_NOT_WORKING);
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
        return this.workdayAPIService.assignPerson(workdayTimeslot.department.id, workdayTimeslot.date, workdayTimeslot.workplace.id, workdayTimeslot.timeslot.id, person.id).pipe(
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

  handleCommentEditRequest(ts: DisplayedWorkdayTimeslot) {
    // generate data for dialog
    const data: EditTextareaDialogData = {
      control: new FormControl(ts.comment),
      label: 'Kommentar zum Timeslot',
      placeholder: 'Blah blah blah',
      hint: 'Dieser Kommentar wird für diesen Timeslot gespeichert.',
    };
    // generate dialogConfig
    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = data;

    // Open dialog to edit comment
    const dialogRef = this.dialog.open(EditTextareaDialogComponent, dialogConfig);
    dialogRef
      .afterClosed()
      .pipe(
        filter((result): result is string => typeof result === 'string'),
        map((comment) => ({
          department_id: ts.department.id,
          workplace_id: ts.workplace.id,
          timeslot_id: ts.timeslot.id,
          date: ts.date,
          comment,
          start_time: ts.start_time,
          end_time: ts.end_time,
          active: true,
        })),
        switchMap((data) => this.workdayAPIService.updateWorkday(data).pipe(catchError((err) => throwError(() => err)))),
        map((resp) => resp.data),
      )
      .subscribe({
        next: (workday) => {
          ts.comment = workday.comment;
        },
      });
  }
}
