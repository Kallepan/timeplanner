import { TestBed } from '@angular/core/testing';

import { DepartmentAPIService } from './department-api.service';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { APIResponse } from '@app/core/interfaces/response';
import { Department, DepartmentWithMetadata } from '../interfaces/department';
import { constants } from '@app/core/constants/constants';
import { AbsenceReponse } from '@app/modules/absency/interfaces/absence';

describe('DepartmentAPIService', () => {
  let service: DepartmentAPIService;
  let httpController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting(), DepartmentAPIService],
    });
    service = TestBed.inject(DepartmentAPIService);
    httpController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  afterEach(() => {
    httpController.verify();
  });

  it('should get absences for a department', () => {
    const mockAbsences: APIResponse<AbsenceReponse[]> = {
      data: [
        {
          person_id: 'person1',
          created_at: new Date(),
          reason: 'sick',
          date: '2021-01-01',
        },
      ],
      status: 200,
      message: 'success',
    };

    const departmentName = 'Engineering';

    service.getAbsencesForDepartment(departmentName, '2021-01-01').subscribe((result) => {
      expect(result).toEqual(mockAbsences);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/department/${departmentName}/absency?date=2021-01-01`);

    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');
  });

  it('should get all departments', () => {
    const mockDepartment: APIResponse<DepartmentWithMetadata[]> = {
      data: [
        {
          id: 'department1',
          name: 'department1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
      ],
      status: 200,
      message: 'success',
    };

    service.getDepartments().subscribe((result) => {
      expect(result).toEqual(mockDepartment);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/department`);
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockDepartment);
  });

  it('should get a department', () => {
    const mockDepartment: APIResponse<DepartmentWithMetadata> = {
      data: {
        id: 'Engineering',
        name: 'Engineering',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      status: 200,
      message: 'success',
    };

    const departmentName = 'Engineering';

    service.getDepartment(departmentName).subscribe((result) => {
      expect(result).toEqual(mockDepartment);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/department/${departmentName}`);
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockDepartment);
  });

  it('should create a department', () => {
    const mockDepartment: APIResponse<DepartmentWithMetadata> = {
      data: {
        name: 'Engineering',
        id: 'Engineering',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      status: 200,
      message: 'success',
    };

    const department: Department = {
      id: 'Engineering',
      name: 'Engineering',
    };

    service.createDepartment(department).subscribe((result) => {
      expect(result).toEqual(mockDepartment);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/department`);
    expect(req.request.method).toEqual('POST');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockDepartment);
  });

  it('should delete a department', () => {
    const mockResponse: APIResponse<null> = {
      data: null,
      status: 200,
      message: 'Department deleted',
    };

    const departmentName = 'Engineering';

    service.deleteDepartment(departmentName).subscribe((result) => {
      expect(result).toEqual(mockResponse);
    });

    const req = httpController.expectOne(`${constants.APIS.PLANNER}/department/${departmentName}`);
    expect(req.request.method).toEqual('DELETE');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockResponse);
  });
});
