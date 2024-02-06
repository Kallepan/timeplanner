import { Injectable, inject } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { messages } from '@app/constants/messages';
import { NotificationService } from '@app/core/services/notification.service';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { EditTextareaDialogComponent, EditTextareaDialogData } from '@app/shared/components/edit-textarea-dialog/edit-textarea-dialog.component';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { catchError, filter, map, of, switchMap, tap, throwError } from 'rxjs';

@Injectable({
  providedIn: null,
})
export class PlannerStateHandlerService {
  // dialog
  private dialog = inject(MatDialog);

  // services
  private notificationService = inject(NotificationService);
  private personAPIService = inject(PersonAPIService);
  private workdayAPIService = inject(WorkdayAPIService);

  assignPersonToTimeslot(person: PersonWithMetadata, workdayTimeslot: WorkdayTimeslot, actionToBeExecutedOnFailedValidation?: () => void) {
    of(workdayTimeslot)
      .pipe(
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
      )
      .subscribe({
        next: () => {
          this.notificationService.infoMessage(messages.PLANNER.TIMESLOT_ASSIGNMENT.SUCCESS);
        },
        // on error do nothing as the error is handled in the http interceptor
      });
  }

  unAssignPersonFromTimeslot(person: PersonWithMetadata, workdayTimeslot: WorkdayTimeslot) {
    return this.workdayAPIService
      .unassignPerson(workdayTimeslot.department.id, workdayTimeslot.date, workdayTimeslot.workplace.id, workdayTimeslot.timeslot.id, person.id)
      .pipe(
        map((resp) => resp.data),
        map(() => {
          workdayTimeslot.person = null;
        }),
      )
      .subscribe({
        next: () => {
          this.notificationService.infoMessage(messages.PLANNER.TIMESLOT_UNASSIGNMENT.SUCCESS);
        },
        // on error do nothing as the error is handled in the http interceptor
      });
  }

  handleCommentEditRequest(ts: DisplayedWorkdayTimeslot) {
    // generate data for dialog
    const data: EditTextareaDialogData = {
      title: 'Kommentar bearbeiten',
      control: new FormControl(ts.comment),
      label: 'Kommentar zum Timeslot',
      placeholder: 'Blah blah blah',
      hint: 'Dieser Kommentar wird für diesen Timeslot gespeichert.',
    };
    // generate dialogConfig
    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = data;
    dialogConfig.enterAnimationDuration = 300;
    dialogConfig.exitAnimationDuration = 300;

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
