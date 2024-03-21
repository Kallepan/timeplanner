import { TestBed } from '@angular/core/testing';

import { ActivatedRoute } from '@angular/router';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';
import { of } from 'rxjs';

describe('ActiveDepartmentHandlerService', () => {
  let service: ActiveDepartmentHandlerService;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(() => {
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      queryParams: of({ department: 'test' }),
    });

    TestBed.configureTestingModule({
      providers: [
        {
          provide: ActivatedRoute,
          useValue: mockActivatedRoute,
        },
      ],
    });
    service = TestBed.inject(ActiveDepartmentHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get department from router query params', () => {
    expect(service.activeDepartment$).toBe('test');
  });

  it('should handle null department', () => {
    service.activeDepartment = null;

    expect(service.activeDepartment$).toBe(null);
  });

  it('should format department', () => {
    service.activeDepartment = 'TEST';

    expect(service.activeDepartment$).toBe('test');
  });

  it('_activeDepartment should be defined', () => {
    expect(service['_activeDepartment']).toBeDefined();
  });
});
