import { TestBed } from '@angular/core/testing';

import { ActivatedRoute } from '@angular/router';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';

describe('ActiveDepartmentHandlerService', () => {
  let service: ActiveDepartmentHandlerService;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(() => {
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      snapshot: {
        queryParams: {
          department: 'test',
        },
      },
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

  it('should handle undefined department', () => {
    service.activeDepartment = undefined;

    expect(service.activeDepartment$).toBe('');
  });
});
