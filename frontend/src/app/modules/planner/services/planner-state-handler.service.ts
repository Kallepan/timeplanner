import { Injectable, inject } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { messages } from '@app/constants/messages';
import { NotificationService } from '@app/core/services/notification.service';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { ConfirmationDialogComponent, ConfirmationDialogComponentData } from '@app/shared/components/confirmation-dialog/confirmation-dialog.component';
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

  assignPersonToTimeslot(person: PersonWithMetadata, workdayTimeslot: WorkdayTimeslot) {
    of(workdayTimeslot)
      .pipe(
        // check if the person is qualified for the workplace
        map((ts) => (person.workplaces ?? []).map((wp) => wp.id).includes(ts.workplace.id)),
        tap((isQualified) => {
          if (!isQualified) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_NOT_QUALIFIED);
        }),
        // check if the person is absent on the day
        switchMap(() =>
          this.personAPIService.getAbsencyForPerson(person.id, workdayTimeslot.date).pipe(
            map((resp) => resp.data),
            map((data) => {
              if (data) return true;
              return false;
            }),
            catchError(() => throwError(() => new Error(messages.GENERAL.HTTP_ERROR.SERVER_ERROR))),
          ),
        ),
        tap((isAbsent) => {
          if (isAbsent) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_ABSENT);
        }),
        // check if the person is already assigned to a timeslot on the same day
        switchMap(() =>
          this.workdayAPIService.getWorkdays(workdayTimeslot.department.id, workdayTimeslot.date).pipe(catchError(() => throwError(() => new Error(messages.GENERAL.HTTP_ERROR.SERVER_ERROR)))),
        ),
        map((resp) => resp.data),
        map((workdays) =>
          // check if the person is already assigned to a timeslot on the same day
          workdays
            .filter((wd) => wd.date === workdayTimeslot.date && !!wd.persons.length)
            .map((wd) => wd.persons.map((p) => p.id))
            .flat()
            .includes(person.id),
        ),
        tap((isAssigned) => {
          if (isAssigned) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_ALREADY_ASSIGNED);
        }),
        // check if person is usually present on the day
        map(() => {
          return (person.weekdays ?? []).map((wd) => wd.id).includes(workdayTimeslot.weekday);
        }),
        switchMap((isPresent) => {
          if (isPresent) return of(true);
          return this.dialog
            .open(ConfirmationDialogComponent, { data: { title: 'Person abwesend', confirmationMessage: 'Person ist an diesem Tag normalerweise abwesend. Trotzdem zuweisen?' } })
            .afterClosed()
            .pipe(map((result) => result === true));
        }),
        tap((isPresent) => {
          if (!isPresent) throw new Error(messages.PLANNER.TIMESLOT_ASSIGNMENT.PERSON_NOT_WORKING);
        }),
        map(() => true),
        catchError((err: Error) => {
          this.notificationService.warnMessage(err.message);
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
          workdayTimeslot.persons = [...workdayTimeslot.persons, person];
          return true;
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
        // update the timeslot in the service
        map((resp) => resp.data),
        // remove person from timeslot
        map(() => {
          workdayTimeslot.persons = workdayTimeslot.persons.filter((p) => p.id !== person.id);
        }),
      )
      .subscribe({
        next: () => {
          this.notificationService.infoMessage(messages.PLANNER.TIMESLOT_UNASSIGNMENT.SUCCESS);
        },
        // on error do nothing as the error is handled in the http interceptor
      });
  }

  handleCommentDeleteRequest(ts: DisplayedWorkdayTimeslot) {
    const data: ConfirmationDialogComponentData = {
      title: 'Kommentar löschen',
      confirmationMessage: 'Sind Sie sicher, dass Sie den Kommentar löschen möchten?',
    };
    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = data;
    dialogConfig.enterAnimationDuration = 300;
    dialogConfig.exitAnimationDuration = 300;

    this.dialog
      .open(ConfirmationDialogComponent, dialogConfig)
      .afterClosed()
      .pipe(
        filter((result) => result !== null && result !== undefined && result === true),
        switchMap(() =>
          this.workdayAPIService
            .updateWorkday({
              department_id: ts.department.id,
              workplace_id: ts.workplace.id,
              timeslot_id: ts.timeslot.id,
              date: ts.date,
              comment: '',
              start_time: ts.start_time,
              end_time: ts.end_time,
              active: true,
            })
            .pipe(
              catchError((err) => throwError(() => err)),
              map((resp) => resp.data),
            ),
        ),
      )
      .subscribe({
        next: (workday) => {
          ts.comment = workday.comment;
        },
      });
  }

  handleCommentEditRequestForManySlots(tss: DisplayedWorkdayTimeslot[]) {
    // generate data for dialog
    const data: EditTextareaDialogData = {
      title: 'Kommentar bearbeiten',
      control: new FormControl(''),
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
      .pipe(filter((result): result is string => typeof result === 'string'))
      .subscribe({
        next: (comment) => {
          tss.forEach((ts) => {
            this.workdayAPIService
              .updateWorkday({
                department_id: ts.department.id,
                workplace_id: ts.workplace.id,
                timeslot_id: ts.timeslot.id,
                date: ts.date,
                comment,
                start_time: ts.start_time,
                end_time: ts.end_time,
                active: true,
              })
              .pipe(
                catchError((err) => throwError(() => err)),
                map((resp) => resp.data),
              )
              .subscribe({
                next: (workday) => {
                  ts.comment = workday.comment;
                },
              });
          });
        },
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
