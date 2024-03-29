import { TestBed } from '@angular/core/testing';

import { WorkdayAPIService } from './workday-api.service';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { APIResponse } from '@app/core/interfaces/response';

import { constants } from '@app/core/constants/constants';
import { provideHttpClient } from '@angular/common/http';
import { AssignPersonToWorkdayTimeslotRequest, UnassignPersonFromWorkdayTimeslotRequest, WorkdayTimeslot } from '../interfaces/workday_timeslot';
import { DepartmentWithMetadata } from '../interfaces/department';
import { WorkplaceWithMetadata } from '../interfaces/workplace';
import { TimeslotWithMetadata } from '../interfaces/timeslot';

const MOCK_WORKDAY: WorkdayTimeslot = {
  department: {
    name: 'department1',
    id: 'department1',
  } as DepartmentWithMetadata,
  date: '2022-01-01',
  workplace: {
    name: 'workplace1',
    id: 'workplace1',
  } as WorkplaceWithMetadata,
  timeslot: {
    name: 'timeslot1',
    id: 'timeslot1',
  } as TimeslotWithMetadata,
  start_time: '09:00',
  end_time: '10:00',
  weekday: 1,
  duration_in_minutes: 60,
  comment: '',

  persons: [],
};

describe('WorkdayAPIService', () => {
  let service: WorkdayAPIService;
  let httpController: HttpTestingController;

  beforeEach(async () => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting(), WorkdayAPIService],
    });
    service = TestBed.inject(WorkdayAPIService);
    httpController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  afterEach(() => {
    httpController.verify();
  });

  it('should fetch workday details', async () => {
    const mockWorkday: APIResponse<WorkdayTimeslot> = {
      data: MOCK_WORKDAY,
      status: 200,
      message: 'success',
    };

    const departmentName = 'department1';
    const date = '2022-01-01';
    const workplace = 'workplace1';
    const timeslot = 'timeslot1';

    service.getDetailWorkday(departmentName, date, workplace, timeslot).subscribe((workday) => {
      expect(workday).toEqual(mockWorkday);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/workday/detail?department=${departmentName}&date=${date}&workplace=${workplace}&timeslot=${timeslot}`);
    req.flush(mockWorkday);

    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');
  });

  it('should fetch workdays', () => {
    const mockWorkdays: APIResponse<WorkdayTimeslot[]> = {
      data: [MOCK_WORKDAY, MOCK_WORKDAY], // replace with your actual mock workdays
      status: 200,
      message: 'success',
    };

    const departmentName = 'department1';
    const date = '2022-01-01';

    service.getWorkdays(departmentName, date).subscribe((workdays) => {
      expect(workdays).toEqual(mockWorkdays);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/workday?department=${departmentName}&date=${date}`);
    req.flush(mockWorkdays);

    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');
  });

  it('should unassign person from workday', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'success',
    };

    const departmentName = 'department1';
    const date = '2022-01-01';
    const workplace = 'workplace1';
    const timeslot = 'timeslot1';
    const personId = 'person1';

    service.unassignPerson(departmentName, date, workplace, timeslot, personId).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/workday/assign`);
    expect(req.request.method).toEqual('DELETE');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    const expectedBody: UnassignPersonFromWorkdayTimeslotRequest = {
      department_id: departmentName,
      date: date,
      workplace_id: workplace,
      timeslot_id: timeslot,
      person_id: personId,
    };
    expect(req.request.body).toEqual(expectedBody);

    req.flush(mockResponse);
  });

  it('should assign person to workday', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'success',
    };

    const departmentName = 'department1';
    const date = '2022-01-01';
    const workplace = 'workplace1';
    const timeslot = 'timeslot1';
    const personId = 'person1';

    service.assignPerson(departmentName, date, workplace, timeslot, personId).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/workday/assign`);
    expect(req.request.method).toEqual('POST');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    const expectedBody: AssignPersonToWorkdayTimeslotRequest = {
      department_id: departmentName,
      date: date,
      workplace_id: workplace,
      timeslot_id: timeslot,
      person_id: personId,
    };
    expect(req.request.body).toEqual(expectedBody);

    req.flush(mockResponse);
  });
});
