import { TestBed } from '@angular/core/testing';

import { TimeslotAPIService } from './timeslot-api.service';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';

describe('TimeslotAPIService', () => {
  let service: TimeslotAPIService;
  let httpController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting(), TimeslotAPIService],
    });
    service = TestBed.inject(TimeslotAPIService);

    httpController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  afterEach(() => {
    httpController.verify();
  });

  it('should get all timeslots', () => {
    const mockTimeslot = {
      data: [
        {
          id: 'timeslot1',
          name: 'timeslot1',
          department_id: 'department1',
          workplace_id: 'workplace1',
          active: true,
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          weekdays: [],
        },
      ],
      status: 200,
      message: 'success',
    };

    service.getTimeslots('department1', 'workplace1').subscribe((result) => {
      expect(result).toEqual(mockTimeslot);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1/timeslot');
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockTimeslot);
  });

  it('should get a timeslot', () => {
    const mockTimeslot = {
      data: {
        id: 'timeslot1',
        name: 'timeslot1',
        department_id: 'department1',
        workplace_id: 'workplace1',
        active: true,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        weekdays: [],
      },
      status: 200,
      message: 'success',
    };

    service.getTimeslot('department1', 'workplace1', 'timeslot1').subscribe((result) => {
      expect(result).toEqual(mockTimeslot);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1/timeslot/timeslot1');
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockTimeslot);
  });

  it('should create a timeslot', () => {
    const mockTimeslot = {
      data: {
        id: 'timeslot1',
        name: 'timeslot1',
        department_id: 'department1',
        workplace_id: 'workplace1',
        active: true,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        weekdays: [],
      },
      status: 200,
      message: 'success',
    };

    service
      .createTimeslot('department1', 'workplace1', {
        id: 'timeslot1',
        name: 'timeslot1',
        active: true,
      })
      .subscribe((result) => {
        expect(result).toEqual(mockTimeslot);
      });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1/timeslot');
    expect(req.request.method).toEqual('POST');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');
    expect(req.request.body).toEqual({
      name: 'timeslot1',
      id: 'timeslot1',
      active: true,
    });
    req.flush(mockTimeslot);
  });

  it('should delete a timeslot', () => {
    service.deleteTimeslot('department1', 'workplace1', 'timeslot1').subscribe((result) => {
      expect(result).toEqual({
        data: null,
        status: 200,
        message: 'success',
      });
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1/timeslot/timeslot1');
    expect(req.request.method).toEqual('DELETE');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush({
      data: null,
      status: 200,
      message: 'success',
    });
  });

  it('should assign a weekday to a timeslot', () => {
    const mockTimeslot = {
      data: {
        name: 'timeslot1',
        id: 'timeslot1',
        department_id: 'department1',
        workplace_id: 'workplace1',
        active: true,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        weekdays: [
          {
            id: 1,
            name: 'Monday',
            start_time: '08:00',
            end_time: '16:00',
          },
        ],
      },
      status: 200,
      message: 'success',
    };

    service.assignTimeslotToWeekday('department1', 'workplace1', 'timeslot1', 1, '08:00', '16:00').subscribe((result) => {
      expect(result).toEqual(mockTimeslot);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1/timeslot/timeslot1/weekday');
    expect(req.request.method).toEqual('POST');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');
    expect(req.request.body).toEqual({
      id: 1,
      start_time: '08:00',
      end_time: '16:00',
    });

    req.flush(mockTimeslot);
  });

  it('should remove a weekday from a timeslot', () => {
    const mockTimeslot = {
      data: {
        name: 'timeslot1',
        id: 'timeslot1',
        department_id: 'department1',
        workplace_id: 'workplace1',
        active: true,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        weekdays: [],
      },
      status: 200,
      message: 'success',
    };

    service.unassignTimeslotFromWeekday('department1', 'workplace1', 'timeslot1', 1).subscribe((result) => {
      expect(result).toEqual(mockTimeslot);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1/timeslot/timeslot1/weekday');
    expect(req.request.method).toEqual('DELETE');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');
    expect(req.request.body).toEqual({
      id: 1,
    });
    req.flush(mockTimeslot);
  });
});
