import { TestBed } from '@angular/core/testing';

import { WorkplaceAPIService } from './workplace-api.service';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';

describe('WorkplaceAPIService', () => {
  let service: WorkplaceAPIService;
  let httpController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting(), WorkplaceAPIService],
    });
    service = TestBed.inject(WorkplaceAPIService);
    httpController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  afterEach(() => {
    httpController.verify();
  });

  it('should get all workplaces', () => {
    const mockWorkplace = {
      data: [
        {
          id: 'workplace1',
          name: 'workplace1',
          department_id: 'department1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
      ],
      status: 200,
      message: 'success',
    };

    service.getWorkplaces('department1').subscribe((result) => {
      expect(result).toEqual(mockWorkplace);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace');
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockWorkplace);
  });

  it('should get a workplace', () => {
    const mockWorkplace = {
      data: {
        name: 'workplace1',

        id: 'workplace1',
        department_id: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      status: 200,
      message: 'success',
    };

    service.getWorkplace('department1', 'workplace1').subscribe((result) => {
      expect(result).toEqual(mockWorkplace);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1');
    expect(req.request.method).toEqual('GET');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockWorkplace);
  });

  it('should create a workplace', () => {
    const mockWorkplace = {
      data: {
        name: 'workplace1',
        id: 'workplace1',
        department_id: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      status: 200,
      message: 'success',
    };

    service
      .createWorkplace('department1', {
        id: 'workplace1',
        name: 'workplace1',
      })
      .subscribe((result) => {
        expect(result).toEqual(mockWorkplace);
      });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace');
    expect(req.request.method).toEqual('POST');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockWorkplace);
  });

  it('should delete a workplace', () => {
    const mockWorkplace = {
      data: {
        name: 'workplace1',
        id: 'workplace1',
        department_id: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      status: 200,
      message: 'success',
    };

    service.deleteWorkplace('department1', 'workplace1').subscribe((result) => {
      expect(result).toEqual(mockWorkplace);
    });

    const req = httpController.expectOne('http://localhost:8080/api/v1/planner/department/department1/workplace/workplace1');
    expect(req.request.method).toEqual('DELETE');
    expect(req.request.withCredentials).toEqual(true);
    expect(req.request.headers.get('Content-Type')).toEqual('application/json');

    req.flush(mockWorkplace);
  });
});
