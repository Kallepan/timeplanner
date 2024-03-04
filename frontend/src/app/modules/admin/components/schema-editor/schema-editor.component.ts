import { Component, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTreeModule } from '@angular/material/tree';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { FlatTreeControl } from '@angular/cdk/tree';
import { DynamicDataSource, DynamicFlatNode } from '../../services/tree-data-source';
import { MatIconModule } from '@angular/material/icon';
import { MatTooltipModule } from '@angular/material/tooltip';
import { EditWeekdaysOfTimeslotComponent, POSSIBLE_WEEKDAYS } from './edit-weekdays-of-timeslot/edit-weekdays-of-timeslot.component';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { TimeslotAPIService } from '@app/shared/services/timeslot-api.service';
import { FormControl, FormGroup } from '@angular/forms';
import { NotificationService } from '@app/core/services/notification.service';
import { messages } from '@app/constants/messages';
import { catchError, filter, map, of, switchMap, tap, throwError } from 'rxjs';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ConfirmationDialogComponent } from '@app/shared/components/confirmation-dialog/confirmation-dialog.component';
import { SchemaEditDialogComponent, SchemaEditDialogData } from './schema-edit-dialog/schema-edit-dialog.component';
import { DepartmentWithMetadata } from '@app/shared/interfaces/department';

@Component({
  selector: 'app-schema-editor',
  standalone: true,
  imports: [CommonModule, MatTreeModule, MatButtonModule, MatProgressBarModule, MatIconModule, MatTooltipModule, EditWeekdaysOfTimeslotComponent],
  templateUrl: './schema-editor.component.html',
  styleUrl: './schema-editor.component.scss',
})
export class SchemaEditorComponent {
  // inject services
  private departmentAPIService = inject(DepartmentAPIService);
  private workplaceAPIService = inject(WorkplaceAPIService);
  private timeslotAPIService = inject(TimeslotAPIService);
  private notificationService = inject(NotificationService);

  // setup tree
  treeControl: FlatTreeControl<DynamicFlatNode>;
  dataSource: DynamicDataSource;

  // dialog
  private matDialog = inject(MatDialog);

  constructor() {
    this.treeControl = new FlatTreeControl<DynamicFlatNode>(this.getLevel, this.isExpandable);
    this.dataSource = new DynamicDataSource(this.treeControl);
  }

  getLevel = (node: DynamicFlatNode): number => node.level;
  isExpandable = (node: DynamicFlatNode): boolean => node.expandable;
  hasChild = (_: number, node: DynamicFlatNode): boolean => node.expandable;

  openElement(node: DynamicFlatNode) {
    if (node.type !== 'timeslot') return;

    this._selectedTimeslotForEditing.set(node.item as TimeslotWithMetadata);
  }
  editElement(node: DynamicFlatNode) {
    // TODO: implement
    console.log(node);
  }
  addElement(node: DynamicFlatNode) {
    of(node.type)
      .pipe(
        map((type) => {
          switch (type) {
            case 'department':
              return this.workplaceAPIService;
            case 'workplace':
              return this.timeslotAPIService;
            default:
              return null;
          }
        }),
        filter((result): result is TimeslotAPIService | WorkplaceAPIService => result !== null),
        switchMap((svc) => {
          const data: SchemaEditDialogData<WorkplaceAPIService | TimeslotAPIService> = {
            id: '',
            name: '',
            idIsEditable: true,
            serviceForValidation: svc,
            departmentID: node.type === 'department' ? (node.item as DepartmentWithMetadata).id : (node.item as WorkplaceWithMetadata).department_id,
            workplaceID: node.type === 'workplace' ? (node.item as WorkplaceWithMetadata).id : undefined,
          };

          const matDialogConfig: MatDialogConfig = new MatDialogConfig();
          matDialogConfig.data = data;
          matDialogConfig.enterAnimationDuration = 300;
          matDialogConfig.exitAnimationDuration = 300;

          return this.matDialog.open(SchemaEditDialogComponent, matDialogConfig).afterClosed();
        }),
        filter((result) => result !== null && result !== undefined),
        switchMap((result) => {
          if (node.type === 'department') {
            return this.workplaceAPIService.createWorkplace((node.item as DepartmentWithMetadata).id, { id: result.id, name: result.name }).pipe(
              map((res) => res.data),
              catchError((err) => throwError(() => err)),
              tap((wp) => {
                // insert workplace after the current department
                const data = this.dataSource.data;
                const nodeIndex = data.findIndex((item) => item === node);
                data.splice(nodeIndex + 1, 0, new DynamicFlatNode(wp, node.level + 1, 'workplace'));
                this.dataSource.data = data;
              }),
            );
          }
          if (node.type === 'workplace') {
            const departmentID = (node.item as WorkplaceWithMetadata).department_id;
            const workplaceID = (node.item as WorkplaceWithMetadata).id;
            return this.timeslotAPIService.createTimeslot(departmentID, workplaceID, { id: result.id, name: result.name, active: true }).pipe(
              map((res) => res.data),
              catchError((err) => throwError(() => err)),
              tap((ts) => {
                // insert timeslot after the current workplace
                const data = this.dataSource.data;
                const nodeIndex = data.findIndex((item) => item === node);
                data.splice(nodeIndex + 1, 0, new DynamicFlatNode(ts, node.level + 1, 'timeslot'));
                this.dataSource.data = data;
              }),
            );
          }
          return of(null);
        }),
      )
      .subscribe({
        next: () => {
          this.notificationService.infoMessage(messages.ADMIN.CREATE_SUCCESSFUL);
        },
        error: () => {
          this.notificationService.warnMessage(messages.ADMIN.CREATE_FAILED);
        },
      });
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  deleteElement(node: DynamicFlatNode) {
    of(node)
      .pipe(
        // open confirmation dialog
        map((node) => node.type),
        switchMap((type) => {
          return this.matDialog
            .open(ConfirmationDialogComponent, {
              data: {
                title: 'Löschen bestätigen',
                confirmationMessage: `Sind Sie sicher, dass Sie ${type === 'timeslot' ? 'den Timeslot' : 'den Arbeitsplatz'} löschen möchten?`,
              },
            })
            .afterClosed();
        }),
        filter((result) => result === true),
        // This is clearly a shitty way to do this
        map(() => node),
        switchMap((node) => {
          if (node.type === 'timeslot') {
            const ts = node.item as TimeslotWithMetadata;
            return this.timeslotAPIService.deleteTimeslot(ts.department_id, ts.workplace_id, ts.id).pipe(
              tap(() => {
                this.dataSource.data = this.dataSource.data.filter((item) => item !== node);
              }),
              catchError((err) => throwError(() => err)),
            );
          }
          if (node.type === 'workplace') {
            const wp = node.item as WorkplaceWithMetadata;
            return this.workplaceAPIService.deleteWorkplace(wp.department_id, wp.id).pipe(
              tap(() => {
                this.dataSource.data = this.dataSource.data.filter((item) => item !== node);
              }),
              catchError((err) => throwError(() => err)),
            );
          }
          return of(null);
        }),
        filter((result) => result !== null),
      )
      .subscribe({
        next: () => {
          this.notificationService.infoMessage(messages.ADMIN.DELETE_SUCCESSFUL);
        },
        error: () => {
          this.notificationService.warnMessage(messages.ADMIN.DELETE_FAILED);
        },
      });
  }

  private _selectedTimeslotForEditing = signal<null | TimeslotWithMetadata>(null);
  get selectedTimeslotForEditing$() {
    return this._selectedTimeslotForEditing();
  }

  // handle child events
  addRequest(event: { control: FormGroup; type: string; timeslot: TimeslotWithMetadata }) {
    const timeslot = event.timeslot;
    const weekdayID = event.control.controls['weekday'].value;
    const startTime = event.control.controls['startTime'].value;
    const endTime = event.control.controls['endTime'].value;
    this.timeslotAPIService.assignTimeslotToWeekday(timeslot.department_id, timeslot.workplace_id, timeslot.id, weekdayID, startTime, endTime).subscribe({
      next: () => {
        const selectedWeekday = POSSIBLE_WEEKDAYS.find((weekday) => weekday.id === weekdayID);
        timeslot.weekdays = (timeslot.weekdays ?? []).concat([{ id: weekdayID, name: selectedWeekday?.name ?? '', start_time: startTime, end_time: endTime }]).sort((a, b) => a.id - b.id);
        this.notificationService.infoMessage(messages.ADMIN.TIMESLOT_WEEKDAY_ASSIGNED);
      },
      error: () => {
        this.notificationService.warnMessage(messages.ADMIN.TIMESLOT_WEEKDAY_ASSIGNMENT_FAILED);
      },
    });
  }

  editRequest(event: { weekdayID: number; startTimeControl: FormControl; endTimeControl: FormControl; timeslot: TimeslotWithMetadata }) {
    const { startTimeControl, endTimeControl, timeslot, weekdayID } = event;
    if (startTimeControl.invalid || endTimeControl.invalid) return;

    this.timeslotAPIService.updateTimeslotOnWeekday(timeslot.department_id, timeslot.workplace_id, timeslot.id, weekdayID, startTimeControl.value, endTimeControl.value).subscribe({
      next: () => {
        const weekday = timeslot.weekdays?.find((weekday) => weekday.id === weekdayID);
        if (weekday) {
          weekday.start_time = startTimeControl.value;
          weekday.end_time = endTimeControl.value;
        }
        this.notificationService.infoMessage(messages.ADMIN.TIMESLOT_WEEKDAY_UPDATE_SUCCESS);
      },
      error: () => {
        this.notificationService.warnMessage(messages.ADMIN.TIMESLOT_WEEKDAY_UPDATE_FAILED);
      },
    });
  }

  removeRequest(event: { id: number; timeslot: TimeslotWithMetadata }) {
    const { id, timeslot } = event;
    this.timeslotAPIService.unassignTimeslotFromWeekday(timeslot.department_id, timeslot.workplace_id, timeslot.id, id).subscribe({
      next: () => {
        timeslot.weekdays = (timeslot.weekdays ?? []).filter((weekday) => weekday.id !== id);
        this.notificationService.infoMessage(messages.ADMIN.TIMESLOT_WEEKDAY_UNASSIGNED);
      },
      error: () => {
        this.notificationService.warnMessage(messages.ADMIN.TIMESLOT_WEEKDAY_UNASSIGNMENT_FAILED);
      },
    });
  }
}
