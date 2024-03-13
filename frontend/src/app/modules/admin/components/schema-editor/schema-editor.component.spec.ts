import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SchemaEditorComponent } from './schema-editor.component';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { TimeslotAPIService } from '@app/shared/services/timeslot-api.service';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { NotificationService } from '@app/core/services/notification.service';
import { of, throwError } from 'rxjs';
import { DynamicFlatNode } from '../../services/tree-data-source';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { MatDialog } from '@angular/material/dialog';
import { messages } from '@app/constants/messages';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';
import { FormControl, FormGroup } from '@angular/forms';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { DepartmentWithMetadata } from '@app/shared/interfaces/department';

describe('SchemaEditorComponent', () => {
  let component: SchemaEditorComponent;
  let fixture: ComponentFixture<SchemaEditorComponent>;
  let mockWorkplaceAPIService: jasmine.SpyObj<WorkplaceAPIService>;
  let mockTimeslotAPIService: jasmine.SpyObj<TimeslotAPIService>;
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockDialog: jasmine.SpyObj<MatDialog>;

  beforeEach(async () => {
    mockWorkplaceAPIService = jasmine.createSpyObj('WorkplaceAPIService', ['getWorkplaces', 'deleteWorkplace', 'createWorkplace']);
    mockTimeslotAPIService = jasmine.createSpyObj('TimeslotAPIService', [
      'getTimeslots',
      'deleteTimeslot',
      'createTimeslot',
      'unassignTimeslotFromWeekday',
      'assignTimeslotToWeekday',
      'updateTimeslotOnWeekday',
    ]);
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['getDepartments']);
    mockDepartmentAPIService.getDepartments.and.returnValue(of({ data: [], status: 200, message: 'success' }));

    mockNotificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);

    mockDialog = jasmine.createSpyObj('MatDialog', ['open']);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockDialog.open.and.returnValue({ afterClosed: () => of(true) } as any);

    await TestBed.configureTestingModule({
      imports: [SchemaEditorComponent],
      providers: [
        { provide: WorkplaceAPIService, useValue: mockWorkplaceAPIService },
        { provide: TimeslotAPIService, useValue: mockTimeslotAPIService },
        { provide: DepartmentAPIService, useValue: mockDepartmentAPIService },
        { provide: NotificationService, useValue: mockNotificationService },
        { provide: MatDialog, useValue: mockDialog },
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(SchemaEditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    expect(component.dataSource).toBeTruthy();
    expect(component.treeControl).toBeTruthy();
  });

  it('should reject timeslot creation', () => {
    const mockWorkplace: WorkplaceWithMetadata = {
      id: 'wp1',
      department_id: 'd1',
      name: 'workplace1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const mockWorkplaceNode = new DynamicFlatNode(mockWorkplace, 1, 'workplace', false);
    mockDialog.open.and.returnValue({
      afterClosed: () => of(null),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    component.addElement(mockWorkplaceNode);

    expect(mockTimeslotAPIService.createTimeslot).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
  });

  it('should reject workplace creation', () => {
    const mockDepartment: DepartmentWithMetadata = {
      id: 'd1',
      name: 'department1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const mockDepartmentNode = new DynamicFlatNode(mockDepartment, 0, 'department', true);
    mockDialog.open.and.returnValue({
      afterClosed: () => of(null),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    component.addElement(mockDepartmentNode);

    expect(mockWorkplaceAPIService.createWorkplace).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
  });

  it('should add a workplace', () => {
    const mockDepartment: DepartmentWithMetadata = {
      id: 'd1',
      name: 'department1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const mockDepartmentNode = new DynamicFlatNode(mockDepartment, 0, 'department', true);
    const mockWorkplace: WorkplaceWithMetadata = {
      id: 'wp1',
      department_id: 'd1',
      name: 'workplace1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const mockWorkplaceNode = new DynamicFlatNode(mockWorkplace, 1, 'workplace', false);

    mockDialog.open.and.returnValue({
      afterClosed: () =>
        of({
          id: 'wp1',
          name: 'workplace1',
        }),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    mockWorkplaceAPIService.createWorkplace.and.returnValue(of({ status: 200, message: 'success', data: mockWorkplace }));

    component.addElement(mockDepartmentNode);

    expect(mockWorkplaceAPIService.createWorkplace).toHaveBeenCalledWith(mockDepartment.id, { id: mockWorkplace.id, name: mockWorkplace.name });
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.CREATE_SUCCESSFUL);
    expect(component.dataSource.data).toContain(mockWorkplaceNode);
  });

  it('should add a timeslot', () => {
    const mockWorkplace: WorkplaceWithMetadata = {
      id: 'wp1',
      department_id: 'd1',
      name: 'workplace1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const mockWorkplaceNode = new DynamicFlatNode(mockWorkplace, 1, 'workplace', false);
    const mockTimeslot: TimeslotWithMetadata = {
      id: 't1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [],
    };
    const mockTimeslotNode = new DynamicFlatNode(mockTimeslot, 2, 'timeslot', false);

    mockDialog.open.and.returnValue({
      afterClosed: () =>
        of({
          id: 't1',
          name: 'timeslot1',
        }),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    mockTimeslotAPIService.createTimeslot.and.returnValue(of({ status: 200, message: 'success', data: mockTimeslot }));

    component.addElement(mockWorkplaceNode);

    expect(mockTimeslotAPIService.createTimeslot).toHaveBeenCalledWith(mockTimeslot.department_id, mockTimeslot.workplace_id, {
      id: mockTimeslot.id,
      name: mockTimeslot.name,
      active: mockTimeslot.active,
    });
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.CREATE_SUCCESSFUL);
    expect(component.dataSource.data).toContain(mockTimeslotNode);
  });

  it('should set timeslot on openElement', () => {
    expect(component.selectedTimeslotForEditing$).toBe(null);
    const mockTimeslot: TimeslotWithMetadata = {
      id: 't1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [],
    };
    const mockNode: DynamicFlatNode = {
      item: mockTimeslot,
      level: 1,
      type: 'timeslot',
      expandable: false,
      isLoading: false,
    };

    component.openElement(mockNode);
    fixture.detectChanges();
    expect(component.selectedTimeslotForEditing$).toBe(mockTimeslot);
  });

  it('should remove a workplace', () => {
    mockWorkplaceAPIService.deleteWorkplace.and.returnValue(of({ status: 200, message: 'success', data: null }));

    const dummyWorkplace: WorkplaceWithMetadata = {
      id: 'wp1',
      department_id: 'd1',
      name: 'workplace1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const dummyWorkplaceNode = new DynamicFlatNode(dummyWorkplace, 1, 'workplace', false);

    component.deleteElement(dummyWorkplaceNode);

    expect(mockWorkplaceAPIService.deleteWorkplace).toHaveBeenCalledWith(dummyWorkplace.department_id, dummyWorkplace.id);
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.DELETE_SUCCESSFUL);
    expect(mockNotificationService.warnMessage).not.toHaveBeenCalled();
  });

  it('should not remove a workplace', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockDialog.open.and.returnValue({ afterClosed: () => of(null) } as any);
    mockWorkplaceAPIService.deleteWorkplace.and.returnValue(of({ status: 200, message: 'success', data: null }));

    const dummyWorkplace: WorkplaceWithMetadata = {
      id: 'wp1',
      department_id: 'd1',
      name: 'workplace1',
      deleted_at: null,
      created_at: new Date(),
      updated_at: new Date(),
    };
    const dummyWorkplaceNode = new DynamicFlatNode(dummyWorkplace, 1, 'workplace', false);

    component.deleteElement(dummyWorkplaceNode);

    expect(mockWorkplaceAPIService.deleteWorkplace).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
    expect(mockNotificationService.warnMessage).not.toHaveBeenCalled();
  });

  it('should remove a timeslot', () => {
    mockTimeslotAPIService.deleteTimeslot.and.returnValue(of({ status: 200, message: 'success', data: null }));

    const dummyTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [],
    };

    const dummyTimeslotNode = new DynamicFlatNode(dummyTimeslot, 2, 'timeslot', false);

    component.deleteElement(dummyTimeslotNode);

    expect(mockTimeslotAPIService.deleteTimeslot).toHaveBeenCalledWith(dummyTimeslot.department_id, dummyTimeslot.workplace_id, dummyTimeslot.id);
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.DELETE_SUCCESSFUL);
    expect(mockNotificationService.warnMessage).not.toHaveBeenCalled();
  });
  it('should handle error in remove a timeslot', () => {
    mockTimeslotAPIService.deleteTimeslot.and.returnValue(throwError(() => new Error('error')));

    const dummyTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [],
    };

    const dummyTimeslotNode = new DynamicFlatNode(dummyTimeslot, 2, 'timeslot', false);

    component.deleteElement(dummyTimeslotNode);

    expect(mockTimeslotAPIService.deleteTimeslot).toHaveBeenCalledWith(dummyTimeslot.department_id, dummyTimeslot.workplace_id, dummyTimeslot.id);
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();
  });

  it('should handle remove Weekday request', () => {
    const beforeRequestTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [
        {
          id: 0,
          name: 'Monday',
          start_time: '08:00',
          end_time: '17:00',
        },
      ],
    };
    const afterRequestTimeslot: TimeslotWithMetadata = {
      ...beforeRequestTimeslot,
      weekdays: [],
    };

    mockTimeslotAPIService.unassignTimeslotFromWeekday.and.returnValue(of({ status: 200, message: 'success', data: afterRequestTimeslot }));

    component.removeWeekdayRequest({ id: 0, timeslot: beforeRequestTimeslot });

    expect(mockTimeslotAPIService.unassignTimeslotFromWeekday).toHaveBeenCalledWith(beforeRequestTimeslot.department_id, beforeRequestTimeslot.workplace_id, beforeRequestTimeslot.id, 0);
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.TIMESLOT_WEEKDAY_UNASSIGNED);
  });

  it('should handle error in remove Weekday request', () => {
    const beforeRequestTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [
        {
          id: 0,
          name: 'Monday',
          start_time: '08:00',
          end_time: '17:00',
        },
      ],
    };

    mockTimeslotAPIService.unassignTimeslotFromWeekday.and.returnValue(throwError(() => new Error('error')));

    component.removeWeekdayRequest({ id: 0, timeslot: beforeRequestTimeslot });

    expect(mockTimeslotAPIService.unassignTimeslotFromWeekday).toHaveBeenCalledWith(beforeRequestTimeslot.department_id, beforeRequestTimeslot.workplace_id, beforeRequestTimeslot.id, 0);
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
    expect(mockNotificationService.warnMessage).toHaveBeenCalledWith(messages.ADMIN.TIMESLOT_WEEKDAY_UNASSIGNMENT_FAILED);
  });

  it('should handle add Weekday request', () => {
    const beforeRequestTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [],
    };
    const afterRequestTimeslot: TimeslotWithMetadata = {
      ...beforeRequestTimeslot,
      weekdays: [
        {
          id: 0,
          name: 'Monday',
          start_time: '08:00',
          end_time: '17:00',
        },
      ],
    };
    const mockForm = new FormGroup({
      startTime: new FormControl('08:00'),
      endTime: new FormControl('16:00'),
      weekday: new FormControl(0),
    });

    mockTimeslotAPIService.assignTimeslotToWeekday.and.returnValue(of({ status: 200, message: 'success', data: afterRequestTimeslot }));

    component.addWeekdayRequest({ control: mockForm, type: 'timeslot', timeslot: beforeRequestTimeslot });

    expect(mockTimeslotAPIService.assignTimeslotToWeekday).toHaveBeenCalledWith(beforeRequestTimeslot.department_id, beforeRequestTimeslot.workplace_id, beforeRequestTimeslot.id, 0, '08:00', '16:00');
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.TIMESLOT_WEEKDAY_ASSIGNED);
  });

  it('should handle error in add Weekday request', () => {
    const beforeRequestTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [],
    };
    const mockForm = new FormGroup({
      startTime: new FormControl('08:00'),
      endTime: new FormControl('16:00'),
      weekday: new FormControl(0),
    });

    mockTimeslotAPIService.assignTimeslotToWeekday.and.returnValue(throwError(() => new Error('error')));

    component.addWeekdayRequest({ control: mockForm, type: 'timeslot', timeslot: beforeRequestTimeslot });

    expect(mockTimeslotAPIService.assignTimeslotToWeekday).toHaveBeenCalledWith(beforeRequestTimeslot.department_id, beforeRequestTimeslot.workplace_id, beforeRequestTimeslot.id, 0, '08:00', '16:00');
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
    expect(mockNotificationService.warnMessage).toHaveBeenCalledWith(messages.ADMIN.TIMESLOT_WEEKDAY_ASSIGNMENT_FAILED);
  });

  it('should handle edit Weekday request', () => {
    const beforeRequestTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [
        {
          id: 0,
          name: 'Monday',
          start_time: '08:00',
          end_time: '17:00',
        },
      ],
    };
    const afterRequestTimeslot: TimeslotWithMetadata = {
      ...beforeRequestTimeslot,
      weekdays: [
        {
          id: 0,
          name: 'Monday',
          start_time: '08:00',
          end_time: '16:00',
        },
      ],
    };
    const mockEndTimeControl = new FormControl('16:00');
    const mockStartTimeControl = new FormControl('08:00');
    mockTimeslotAPIService.updateTimeslotOnWeekday.and.returnValue(of({ status: 200, message: 'success', data: afterRequestTimeslot }));

    component.editWeekdayRequest({ startTimeControl: mockStartTimeControl, endTimeControl: mockEndTimeControl, timeslot: beforeRequestTimeslot, weekdayID: 0 });

    expect(mockTimeslotAPIService.updateTimeslotOnWeekday).toHaveBeenCalledWith(beforeRequestTimeslot.department_id, beforeRequestTimeslot.workplace_id, beforeRequestTimeslot.id, 0, '08:00', '16:00');
    expect(mockNotificationService.infoMessage).toHaveBeenCalledWith(messages.ADMIN.TIMESLOT_WEEKDAY_UPDATE_SUCCESS);
    expect(mockNotificationService.warnMessage).not.toHaveBeenCalled();
  });

  it('should handle error in edit Weekday request', () => {
    const beforeRequestTimeslot: TimeslotWithMetadata = {
      id: 'ts1',
      name: 'timeslot1',
      department_id: 'd1',
      workplace_id: 'wp1',
      active: true,
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
      weekdays: [
        {
          id: 0,
          name: 'Monday',
          start_time: '08:00',
          end_time: '17:00',
        },
      ],
    };
    const mockEndTimeControl = new FormControl('16:00');
    const mockStartTimeControl = new FormControl('08:00');
    mockTimeslotAPIService.updateTimeslotOnWeekday.and.returnValue(throwError(() => new Error('error')));

    component.editWeekdayRequest({ startTimeControl: mockStartTimeControl, endTimeControl: mockEndTimeControl, timeslot: beforeRequestTimeslot, weekdayID: 0 });

    expect(mockTimeslotAPIService.updateTimeslotOnWeekday).toHaveBeenCalledWith(beforeRequestTimeslot.department_id, beforeRequestTimeslot.workplace_id, beforeRequestTimeslot.id, 0, '08:00', '16:00');
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
    expect(mockNotificationService.warnMessage).toHaveBeenCalledWith(messages.ADMIN.TIMESLOT_WEEKDAY_UPDATE_FAILED);
  });
});
