import { TestBed } from '@angular/core/testing';

import { ViewerStateHandlerService } from './viewer-state-handler.service';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { of, throwError } from 'rxjs';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { DepartmentWithMetadata } from '@app/shared/interfaces/department';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';

describe('ViewerStateHandlerService', () => {
  let service: ViewerStateHandlerService;
  let mockWorkdayAPIService: jasmine.SpyObj<WorkdayAPIService>;

  beforeEach(() => {
    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays']);

    TestBed.configureTestingModule({
      providers: [
        {
          provide: WorkdayAPIService,
          useValue: mockWorkdayAPIService,
        },
        ViewerStateHandlerService,
      ],
    });
    service = TestBed.inject(ViewerStateHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should return an array of workdays for the active week', (done) => {
    // set the active week
    const mockActiveWeek = {
      department: 'department1',
      dates: [new Date(2022, 1, 1), new Date(2022, 1, 2)],
    };

    const mockDepartment: DepartmentWithMetadata = {
      id: 'dep',
      name: 'department1',
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    };

    const mockWorkplace: WorkplaceWithMetadata = {
      id: 'workplace1',
      name: 'workplace1',
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    };

    const mockTimeslot: TimeslotWithMetadata = {
      name: 'timeslot1',
      active: true,
      department_name: 'department1',
      workplace_name: 'workplace1',
      weekdays: [
        {
          id: 'weekday1',
          name: 'Monday',
          start_time: '08:00:00',
          end_time: '16:00:00',
        },
      ],
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    };

    // create a mock workday
    const mockWorkdays: WorkdayTimeslot[] = [
      {
        date: '2022-02-01',
        department: mockDepartment,
        workplace: mockWorkplace,
        timeslot: mockTimeslot,
        start_time: '08:00:00',
        end_time: '16:00:00',
        person: null,
        weekday: 'MON',
      },
      {
        date: '2022-02-01',
        department: mockDepartment,
        workplace: mockWorkplace,
        timeslot: mockTimeslot,
        start_time: '08:00:00',
        end_time: '16:00:00',
        person: null,
        weekday: 'MON',
      },
    ];

    mockWorkdayAPIService.getWorkdays.and.returnValue(of({ data: mockWorkdays, status: 200, message: 'Success' }));

    service.activeWorkdays$.subscribe((workdays) => {
      // expect the workdays to be 14 since we return two workdays for each day in the active week
      expect(workdays.length).toEqual(mockWorkdays.length * 7);
      expect(workdays[0]).toEqual(mockWorkdays[0]);
      expect(workdays[1]).toEqual(mockWorkdays[1]);
      done();
    });

    service.setActiveView(mockActiveWeek.department, mockActiveWeek.dates[0]);
  });

  it('should return an empty array if the API call fails', (done) => {
    // set the active week
    const mockActiveWeek = {
      department: 'department1',
      dates: [new Date(2022, 1, 1), new Date(2022, 1, 2)],
    };

    mockWorkdayAPIService.getWorkdays.and.returnValue(throwError(() => null));

    service.activeWorkdays$.subscribe((workdays) => {
      expect(workdays).toEqual([]);
      done();
    });

    service.setActiveView(mockActiveWeek.department, mockActiveWeek.dates[0]);
  });
});
