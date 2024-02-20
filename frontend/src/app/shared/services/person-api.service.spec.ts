import { TestBed } from '@angular/core/testing';

import { PersonAPIService } from './person-api.service';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { APIResponse } from '@app/core/interfaces/response';
import { PersonWithMetadata, CreatePerson } from '../interfaces/person';
import { constants } from '@app/constants/constants';
import { AbsenceReponse } from '@app/modules/absency/interfaces/absence';

describe('PersonAPIService', () => {
  let service: PersonAPIService;
  let httpController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting(), PersonAPIService],
    });
    service = TestBed.inject(PersonAPIService);

    httpController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  afterEach(() => {
    httpController.verify();
  });

  it('should fetch person details', () => {
    const mockPerson: APIResponse<PersonWithMetadata> = {
      data: {
        id: '1',
        first_name: 'John',
        last_name: 'Doe',
        email: 'john.doe@example.com',
        active: true,
        working_hours: 8,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        workplaces: [],
        departments: [],
        weekdays: [],
      },

      status: 200,
      message: 'success',
    };

    const personId = '1';

    service.getPerson(personId).subscribe((person) => {
      expect(person).toEqual(mockPerson);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personId}`);
    expect(req.request.method).toEqual('GET');
    req.flush(mockPerson);
  });

  it('should fetch persons', () => {
    const mockPerson: APIResponse<PersonWithMetadata[]> = {
      data: [
        {
          id: '1',
          first_name: 'John',
          last_name: 'Doe',
          email: 'test@example.com',
          active: true,
          working_hours: 8,
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          workplaces: [],
          departments: [],
          weekdays: [],
        },
      ],
      status: 200,
      message: 'success',
    };

    const departmentName = 'department1';

    service.getPersons(departmentName).subscribe((person) => {
      expect(person).toEqual(mockPerson);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person?department=${departmentName}`);
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    const expectedParams = { department: departmentName };
    expect(req.request.params.keys()).toEqual(Object.keys(expectedParams));
    expect(req.request.params.get('department')).toEqual(departmentName);

    req.flush(mockPerson);
  });

  it('should create a person', () => {
    const mockPerson: APIResponse<PersonWithMetadata> = {
      data: {
        id: '1',
        first_name: 'John',
        last_name: 'Doe',
        email: 'john.doe@example.com',
        active: true,
        working_hours: 8,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        workplaces: [],
        departments: [],
        weekdays: [],
      },
      status: 201,
      message: 'created',
    };

    const newPerson: CreatePerson = {
      first_name: 'John',
      last_name: 'Doe',
      email: 'john.doe@example.com',
      active: true,
      working_hours: 8,
    };

    service.createPerson(newPerson).subscribe((person) => {
      expect(person).toEqual(mockPerson);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person`);
    expect(req.request.method).toEqual('POST');
    req.flush(mockPerson);
  });

  it('should update a person', () => {
    const mockPerson: APIResponse<PersonWithMetadata> = {
      data: {
        id: '1',
        first_name: 'John',
        last_name: 'Doe',
        email: 'john.doe@example.com',
        active: true,
        working_hours: 8,
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
        workplaces: [],
        departments: [],
        weekdays: [],
      },
      status: 200,
      message: 'updated',
    };

    const updatedPerson: CreatePerson = {
      first_name: 'John',
      last_name: 'Doe',
      email: 'john.doe@example.com',
      active: true,
      working_hours: 8,
    };
    const id = '1';

    service.updatePerson(updatedPerson, id).subscribe((person) => {
      expect(person).toEqual(mockPerson);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${id}`);
    expect(req.request.method).toEqual('PUT');
    req.flush(mockPerson);
  });

  it('should remove a department from a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Department removed',
    };

    const departmentName = 'Engineering';
    const personID = '1';

    service.removeDepartmentFromPerson(departmentName, personID).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/department/${departmentName}`);
    expect(req.request.method).toEqual('DELETE');
    req.flush(mockResponse);
  });

  it('should add a department to a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Department added',
    };

    const departmentName = 'Engineering';
    const personID = '1';

    service.addDepartmentToPerson(departmentName, personID).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/department`);
    expect(req.request.method).toEqual('POST');
    req.flush(mockResponse);
  });

  it('should add a workplace to a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Workplace added',
    };

    const workplaceName = 'Engineering';
    const departmentName = 'Test';
    const personID = '1';

    service.addWorkplaceToPerson(departmentName, workplaceName, personID).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/workplace`);
    expect(req.request.method).toEqual('POST');
    req.flush(mockResponse);
  });

  it('should remove a workplace from a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Workplace removed',
    };

    const workplaceName = 'Engineering';
    const departmentName = 'Test';
    const personID = '1';

    service.removeWorkplaceFromPerson(departmentName, workplaceName, personID).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/workplace`);
    expect(req.request.method).toEqual('DELETE');
    req.flush(mockResponse);
  });

  it('should add weekday to a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Weekday added',
    };

    const weekdayID = '1';
    const personID = '1';

    service.addWeekdayToPerson(weekdayID, personID).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/weekday`);
    expect(req.request.method).toEqual('POST');
    req.flush(mockResponse);
  });

  it('should remove weekday from a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Weekday removed',
    };

    const weekdayID = '1';
    const personID = '1';

    service.removeWeekdayFromPerson(weekdayID, personID).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/weekday/${weekdayID}`);

    expect(req.request.method).toEqual('DELETE');
    req.flush(mockResponse);
  });

  it('should add an absence to a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Absence added',
    };

    const personID = '1';
    const date = '2022-01-01';
    const reason = 'Sick leave';

    service.addAbsencyToPerson(personID, date, reason).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/absency`);
    expect(req.request.method).toEqual('POST');
    req.flush(mockResponse);
  });

  it('should get an absence for a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Absence retrieved',
    };

    const personID = '1';
    const date = '2022-01-01';

    service.getAbsencyForPerson(personID, date).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/absency?date=${date}`);
    expect(req.request.method).toEqual('GET');
    req.flush(mockResponse);
  });

  it('should handle a valid absence response', () => {
    const mockResponse: APIResponse<AbsenceReponse> = {
      data: {
        date: '2022-01-01',
        reason: 'Sick leave',
        person_id: '1',
        created_at: new Date(),
      },
      status: 200,
      message: 'Absence retrieved',
    };

    const personID = '1';
    const date = '2022-01-01';

    service.getAbsencyForPerson(personID, date).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/absency?date=${date}`);
    expect(req.request.method).toEqual('GET');
    req.flush(mockResponse);
  });

  it('should fetch a range of absences for a person', () => {
    const mockResponses: APIResponse<AbsenceReponse[]> = {
      data: [
        {
          date: '2022-01-01',
          reason: 'Sick leave',
          person_id: '1',
          created_at: new Date(),
        },
        {
          date: '2022-01-02',
          reason: 'Sick leave',
          person_id: '1',
          created_at: new Date(),
        },
      ],
      status: 200,
      message: 'Absences retrieved',
    };

    const personID = '1';
    const startDate = '2022-01-01';
    const endDate = '2022-01-02';

    service.getAbsencyForPersonInRange(personID, startDate, endDate).subscribe((response) => {
      expect(response).toEqual(mockResponses);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/absency?start_date=${startDate}&end_date=${endDate}`);
    expect(req.request.method).toEqual('GET');
    req.flush(mockResponses);
  });

  it('should remove an absence from a person', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Absence removed',
    };

    const personID = '1';
    const date = '2022-01-01';

    service.removeAbsencyFromPerson(personID, date).subscribe((response) => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/person/${personID}/absency/${date}`);
    expect(req.request.method).toEqual('DELETE');
    req.flush(mockResponse);
  });
});
