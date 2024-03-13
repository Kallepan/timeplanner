import { TestBed } from '@angular/core/testing';

import { PlannerStateHandlerService } from './planner-state-handler.service';
import { MatDialog } from '@angular/material/dialog';
import { NotificationService } from '@app/core/services/notification.service';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { formatDateToDateString } from '@app/shared/functions/format-date-to-string.function';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { map, of } from 'rxjs';
import { AbsenceReponse } from '@app/modules/absency/interfaces/absence';
import { APIResponse } from '@app/core/interfaces/response';
import { dateToWeekdayID } from '@app/shared/functions/date-to-weekday-id.function';

const mockWorkdayTimeslot: WorkdayTimeslot = {
  department: {
    id: 'test',
    name: 'test',
    created_at: new Date(),
    updated_at: new Date(),
    deleted_at: null,
  },
  workplace: {
    id: 'test',
    name: 'test',
    department_id: 'test',
    created_at: new Date(),
    updated_at: new Date(),
    deleted_at: null,
  },
  timeslot: {
    department_id: 'test',
    workplace_id: 'test',
    active: true,
    weekdays: [
      {
        id: dateToWeekdayID(new Date()),
        name: new Date().toLocaleDateString('en-US', { weekday: 'long' }),
        start_time: '08:00',
        end_time: '16:00',
      },
    ],
    id: 'test',
    name: 'test',

    created_at: new Date(),
    updated_at: new Date(),
    deleted_at: null,
  },
  date: formatDateToDateString(new Date()),
  start_time: '08:00',
  end_time: '16:00',
  weekday: dateToWeekdayID(new Date()),
  duration_in_minutes: 480,
  comment: 'test',
  persons: [],
};

describe('PlannerStateHandlerService', () => {
  let service: PlannerStateHandlerService;

  // mocks
  let mockDialog: jasmine.SpyObj<MatDialog>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockWorkdayAPIService: jasmine.SpyObj<WorkdayAPIService>;

  beforeEach(() => {
    mockDialog = jasmine.createSpyObj('MatDialog', ['open']);
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['getAbsencyForPerson']);
    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays', 'updateWorkday', 'unassignPerson', 'assignPerson']);

    TestBed.configureTestingModule({
      providers: [
        PlannerStateHandlerService,
        { provide: MatDialog, useValue: mockDialog },
        { provide: NotificationService, useValue: mockNotificationService },
        { provide: PersonAPIService, useValue: mockPersonAPIService },
        { provide: WorkdayAPIService, useValue: mockWorkdayAPIService },
      ],
    });
    service = TestBed.inject(PlannerStateHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  const tests: {
    name: string;
    person: PersonWithMetadata;
    workdayTimeslot: DisplayedWorkdayTimeslot;
    mockPersonAPIServiceResponse: APIResponse<AbsenceReponse | null>;
    mockWorkdayAPIServiceResponse: APIResponse<WorkdayTimeslot[]>;
    mockAssignPersonResponse: APIResponse<null>;
    mockDialogReturnValue: {
      afterClosed: () => void;
    };

    shouldCallDialogTimes?: number; // TODO
    shouldOpenConfirmDialog?: boolean;
    shouldAssignPerson?: boolean;
    shouldFetchAbsency?: boolean;
  }[] = [
    {
      name: 'should not assign person if it (gender neutral :D) is already assigned on another timeslot',
      person: {
        id: 'test',
        workplaces: [{ id: 'test' }],
        weekdays: [{ id: dateToWeekdayID(new Date()) }],
      } as PersonWithMetadata,
      workdayTimeslot: { ...mockWorkdayTimeslot, gridColumn: 1, colorForLightMode: 'red', colorForDarkMode: 'blue', validTime: true },
      mockPersonAPIServiceResponse: {
        data: null,
        message: 'test',
        status: 200,
      },
      mockWorkdayAPIServiceResponse: {
        data: [mockWorkdayTimeslot, { ...mockWorkdayTimeslot, persons: [{ id: 'test' } as PersonWithMetadata] }],
        message: 'test',
        status: 200,
      },
      mockAssignPersonResponse: {
        message: 'ok',
        data: null,
        status: 200,
      },
      mockDialogReturnValue: {
        afterClosed: () => {
          return of(true);
        },
      },
      shouldOpenConfirmDialog: false,
      shouldAssignPerson: false,
      shouldFetchAbsency: true,
    },
    {
      name: 'should assign the person to the timeslot if the person is not absent, the timeslot is empty, and the person is usually present on the day',
      person: {
        id: 'test',
        workplaces: [{ id: 'test' }],
        weekdays: [{ id: dateToWeekdayID(new Date()) }],
      } as PersonWithMetadata,
      workdayTimeslot: { ...mockWorkdayTimeslot, gridColumn: 1, colorForLightMode: 'red', colorForDarkMode: 'blue', validTime: true },
      mockPersonAPIServiceResponse: {
        data: null,
        message: 'test',
        status: 200,
      },
      mockWorkdayAPIServiceResponse: {
        data: [mockWorkdayTimeslot],
        message: 'test',
        status: 200,
      },
      mockAssignPersonResponse: {
        message: 'ok',
        data: null,
        status: 200,
      },
      mockDialogReturnValue: {
        afterClosed: () => {
          return of(true);
        },
      },
      shouldOpenConfirmDialog: false,
      shouldAssignPerson: true,
      shouldFetchAbsency: true,
    },
    {
      name: 'should not assign the person to the timeslot if the person is absent',
      person: {
        id: 'test',
        workplaces: [{ id: 'test' }],
        weekdays: [{ id: dateToWeekdayID(new Date()) }],
      } as PersonWithMetadata,
      workdayTimeslot: { ...mockWorkdayTimeslot, gridColumn: 1, colorForLightMode: 'red', colorForDarkMode: 'blue', validTime: true },
      mockPersonAPIServiceResponse: {
        data: {
          person_id: 'test',
          reason: 'test',
          date: formatDateToDateString(new Date()),
          created_at: new Date(),
        },
        message: 'test',
        status: 200,
      },
      mockWorkdayAPIServiceResponse: {
        data: [mockWorkdayTimeslot],
        message: 'test',
        status: 200,
      },
      mockAssignPersonResponse: {
        message: 'ok',
        data: null,
        status: 400,
      },
      mockDialogReturnValue: {
        afterClosed: () => {
          return of(true);
        },
      },
      shouldOpenConfirmDialog: false,
      shouldAssignPerson: false,
      shouldFetchAbsency: true,
    },
    {
      name: 'should open confirmation dialog if person is not usually present on the day and assign person on confirmation',
      person: {
        id: 'test',
        workplaces: [{ id: 'test' }],
        weekdays: [],
      } as unknown as PersonWithMetadata,
      workdayTimeslot: { ...mockWorkdayTimeslot, gridColumn: 1, colorForLightMode: 'red', colorForDarkMode: 'blue', validTime: true },
      mockPersonAPIServiceResponse: {
        data: null,
        message: 'test',
        status: 200,
      },
      mockWorkdayAPIServiceResponse: {
        data: [mockWorkdayTimeslot],
        message: 'test',
        status: 200,
      },
      mockAssignPersonResponse: {
        message: 'ok',
        data: null,
        status: 200,
      },
      mockDialogReturnValue: {
        afterClosed: () => {
          return of(true);
        },
      },
      shouldOpenConfirmDialog: true,
      shouldAssignPerson: true,
      shouldFetchAbsency: true,
    },
    {
      name: 'should open confirmation dialog if person is not usually present on the day and not assign person on confirmation',
      person: {
        id: 'test',
        workplaces: [{ id: 'test' }],
        weekdays: [],
      } as unknown as PersonWithMetadata,
      workdayTimeslot: { ...mockWorkdayTimeslot, gridColumn: 1, colorForLightMode: 'red', colorForDarkMode: 'blue', validTime: true },
      mockPersonAPIServiceResponse: {
        data: null,
        message: 'test',
        status: 200,
      },
      mockWorkdayAPIServiceResponse: {
        data: [mockWorkdayTimeslot],
        message: 'test',
        status: 200,
      },
      mockAssignPersonResponse: {
        message: 'ok',
        data: null,
        status: 200,
      },
      mockDialogReturnValue: {
        afterClosed: () => {
          return of(false);
        },
      },
      shouldOpenConfirmDialog: true,
      shouldAssignPerson: false,
      shouldFetchAbsency: true,
    },
  ];
  tests.forEach((t) => {
    it(t.name, () => {
      // reset the mocks
      mockPersonAPIService.getAbsencyForPerson.calls.reset();
      mockWorkdayAPIService.assignPerson.calls.reset();
      mockDialog.open.calls.reset();

      // setup the mocks
      mockPersonAPIService.getAbsencyForPerson.and.returnValue(of(t.mockPersonAPIServiceResponse));
      mockWorkdayAPIService.getWorkdays.and.returnValue(of(t.mockWorkdayAPIServiceResponse));
      mockWorkdayAPIService.assignPerson.and.returnValue(of(t.mockAssignPersonResponse));
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      mockDialog.open.and.returnValue(t.mockDialogReturnValue as any);

      // call the method to test
      service.assignPersonToTimeslot(t.person, t.workdayTimeslot);

      // check if the methods were called
      if (t.shouldOpenConfirmDialog) {
        expect(mockDialog.open).toHaveBeenCalled();
      } else {
        expect(mockDialog.open).not.toHaveBeenCalled();
      }

      if (t.shouldAssignPerson) {
        expect(mockWorkdayAPIService.assignPerson).toHaveBeenCalledWith(
          t.workdayTimeslot.department.id,
          t.workdayTimeslot.date,
          t.workdayTimeslot.workplace.id,
          t.workdayTimeslot.timeslot.id,
          t.person.id,
        );
        expect(t.workdayTimeslot.persons.length).toBe(1);
        expect(mockNotificationService.infoMessage).toHaveBeenCalled();
      } else {
        expect(mockWorkdayAPIService.assignPerson).not.toHaveBeenCalled();
        expect(t.workdayTimeslot.persons.length).toBe(0);
      }

      if (t.shouldFetchAbsency) {
        expect(mockPersonAPIService.getAbsencyForPerson).toHaveBeenCalled();
      } else {
        expect(mockPersonAPIService.getAbsencyForPerson).not.toHaveBeenCalled();
      }
    });
  });

  it('should handle erroneous unassignment from unAssignPersonFromTimeslot', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.unassignPerson.and.returnValue(
      of().pipe(
        map(() => {
          throw new Error('test');
        }),
      ),
    );

    service.unAssignPersonFromTimeslot(ts.persons[0], ts);

    expect(mockDialog.open).not.toHaveBeenCalled();
    expect(mockWorkdayAPIService.unassignPerson).toHaveBeenCalled();

    expect(ts.persons.length).toBe(1);
  });

  it('should unassign one person from the timeslot if multiple persons are assigned using unAssignPersonFromTimeslot', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
      {
        id: 'test2',
      } as PersonWithMetadata,
    ];

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.unassignPerson.and.returnValue(of({ message: 'test', status: 200, data: { comment: '' } as any }));

    service.unAssignPersonFromTimeslot(ts.persons[0], ts);

    expect(mockDialog.open).not.toHaveBeenCalled();
    expect(mockWorkdayAPIService.unassignPerson).toHaveBeenCalled();
    expect(ts.persons.length).toBe(1);
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should unassign the person from the timeslot using unAssignPersonFromTimeslot', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.unassignPerson.and.returnValue(of({ message: 'test', status: 200, data: { comment: '' } as any }));

    service.unAssignPersonFromTimeslot(ts.persons[0], ts);

    expect(mockDialog.open).not.toHaveBeenCalled();
    expect(mockWorkdayAPIService.unassignPerson).toHaveBeenCalled();
    expect(ts.persons.length).toBe(0);
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should delete the comment from handleCommentDeleteRequest', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    mockDialog.open.and.returnValue({
      afterClosed: () => {
        return of(true);
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.updateWorkday.and.returnValue(of({ message: 'test', status: 200, data: { comment: '' } as any }));

    service.handleCommentDeleteRequest(ts);

    expect(mockWorkdayAPIService.updateWorkday).toHaveBeenCalled();
    expect(ts.comment).toBe('');
  });
  it('should handle Error from handleCommentDeleteRequest', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    mockDialog.open.and.returnValue({
      afterClosed: () => {
        return of(true);
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.updateWorkday.and.returnValue(
      of().pipe(
        map(() => {
          throw new Error('test');
        }),
      ),
    );

    service.handleCommentDeleteRequest(ts);

    // fetch confirm dialog button
    window.document.querySelector('button')?.click();

    expect(mockWorkdayAPIService.updateWorkday).toHaveBeenCalled();
    expect(ts.comment).toBe('test');
  });

  it('should update comment from handleCommentEditRequest', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    mockDialog.open.and.returnValue({
      afterClosed: () => {
        return of('new_comment');
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.updateWorkday.and.returnValue(of({ message: 'test', status: 200, data: { comment: 'new_comment' } as any }));

    service.handleCommentEditRequest(ts);

    expect(mockDialog.open).toHaveBeenCalled();
    expect(mockWorkdayAPIService.updateWorkday).toHaveBeenCalled();
    expect(ts.comment).toBe('new_comment');
  });
  it('should not update comment from handleCommentEditRequest if dialog is closed without a value', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    mockDialog.open.and.returnValue({
      afterClosed: () => {
        return of(undefined);
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    service.handleCommentEditRequest(ts);

    expect(mockDialog.open).toHaveBeenCalled();
    expect(mockWorkdayAPIService.updateWorkday).not.toHaveBeenCalled();
    expect(ts.comment).toBe('test');
  });
  it('should handle Error from handleCommentEditRequest', () => {
    const ts: DisplayedWorkdayTimeslot = {
      ...mockWorkdayTimeslot,
      // Extra Values
      gridColumn: 1,
      colorForLightMode: 'red',
      colorForDarkMode: 'blue',
      validTime: true,
    };

    ts.persons = [
      {
        id: 'test',
      } as PersonWithMetadata,
    ];

    mockDialog.open.and.returnValue({
      afterClosed: () => {
        return of('new_comment');
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockWorkdayAPIService.updateWorkday.and.returnValue(
      of().pipe(
        map(() => {
          throw new Error('test');
        }),
      ),
    );

    service.handleCommentEditRequest(ts);

    expect(mockDialog.open).toHaveBeenCalled();
    expect(mockWorkdayAPIService.updateWorkday).toHaveBeenCalled();
    expect(ts.comment).toBe('test');
  });
});
