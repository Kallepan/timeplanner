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
    created_at: new Date(),
    updated_at: new Date(),
    deleted_at: null,
  },
  timeslot: {
    department_name: 'test',
    workplace_name: 'test',
    active: true,
    weekdays: [
      {
        // Three letter format of the day, e.g. MON, TUE, WED, THU, FRI, SAT, SUN
        id: new Date().toLocaleDateString('en-US', { weekday: 'short' }).toUpperCase(),
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
  weekday: new Date().toLocaleDateString('en-US', { weekday: 'short' }).toUpperCase(),
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
    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays', 'updateWorkday', 'unassignPerson']);

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

  it('should assign the person to the timeslot using assignPersonToTimeslot', () => {
    expect(false).toBeTrue();
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
