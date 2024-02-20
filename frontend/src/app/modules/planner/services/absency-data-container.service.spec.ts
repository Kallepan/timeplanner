import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AbsencyDataContainerService } from './absency-data-container.service';
import { ActivatedRoute } from '@angular/router';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { Component } from '@angular/core';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { of } from 'rxjs';

describe('AbsencyDataContainerService', () => {
  let service: AbsencyDataContainerService;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;

  beforeEach(() => {
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['getAbsencesForDepartment']);

    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      snapshot: {
        queryParams: {
          department: 'test',
        },
      },
    });

    TestBed.configureTestingModule({
      providers: [AbsencyDataContainerService, { provide: ActivatedRoute, useValue: mockActivatedRoute }, { provide: DepartmentAPIService, useValue: mockDepartmentAPIService }],
    });
    service = TestBed.inject(AbsencyDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should have a default absencesGroupedByWeekday', () => {
    expect(service.absencesGroupedByWeekday$).toEqual([]);
  });
});

@Component({
  template: `{{ absencesGroupedByWeekday$ }}`,
})
class AbsencyDataContainerServiceTestComponent {
  private absencyDataContainerService = TestBed.inject(AbsencyDataContainerService);
  absencesGroupedByWeekday$ = this.absencyDataContainerService.absencesGroupedByWeekday$;
}
describe('AbsencyDataContainerService with Component', () => {
  let fixture: ComponentFixture<AbsencyDataContainerServiceTestComponent>;
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  let activeWeekHandlerService: ActiveWeekHandlerService;
  let absencyDataContainerService: AbsencyDataContainerService;

  beforeEach(() => {
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['getAbsencesForDepartment']);
    mockDepartmentAPIService.getAbsencesForDepartment.and.returnValue(
      of({
        data: [],
        message: 'test',
        status: 200,
      }),
    );
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      snapshot: {
        queryParams: {
          department: 'test',
        },
      },
    });

    TestBed.configureTestingModule({
      declarations: [AbsencyDataContainerServiceTestComponent],
      providers: [AbsencyDataContainerService, { provide: ActivatedRoute, useValue: mockActivatedRoute }, { provide: DepartmentAPIService, useValue: mockDepartmentAPIService }],
    });

    fixture = TestBed.createComponent(AbsencyDataContainerServiceTestComponent);
    absencyDataContainerService = TestBed.inject(AbsencyDataContainerService);
    activeWeekHandlerService = TestBed.inject(ActiveWeekHandlerService);
  });

  it('should be created', () => {
    expect(fixture).toBeTruthy();
    expect(absencyDataContainerService).toBeTruthy();
  });

  it('should have a default absencesGroupedByWeekday', () => {
    expect(fixture.componentInstance.absencesGroupedByWeekday$).toEqual([]);
  });

  it('should initialize correctly', async () => {
    expect(absencyDataContainerService.loading$).toBeTruthy();
    fixture.detectChanges();

    activeWeekHandlerService.activeWeekByDate = new Date();
    expect(mockDepartmentAPIService.getAbsencesForDepartment).toHaveBeenCalled();
    fixture.detectChanges();

    expect(absencyDataContainerService.absencesGroupedByWeekday$.length).toBe(7);
    expect(absencyDataContainerService.loading$).toBeFalsy();
  });

  it('should have a default loading state', () => {
    expect(absencyDataContainerService.loading$).toBeTruthy();
  });

  it('should have a default absencesGroupedByWeekday', () => {
    expect(absencyDataContainerService.absencesGroupedByWeekday$).toEqual([]);
  });

  it('should load absences from the API', () => {
    mockDepartmentAPIService.getAbsencesForDepartment.and.returnValue(
      of({
        data: [
          {
            person_id: '1',
            reason: 'test',
            date: '2021-01-01',
            created_at: new Date(),
          },
        ],
        message: 'test',
        status: 200,
      }),
    );

    const date = new Date();
    activeWeekHandlerService.activeWeekByDate = date;
    fixture.detectChanges();

    expect(mockDepartmentAPIService.getAbsencesForDepartment).toHaveBeenCalled();

    // 7 times for each day of the week
    expect(absencyDataContainerService.absencesGroupedByWeekday$.length).toBe(7);
    expect(absencyDataContainerService.absencesGroupedByWeekday$[0]).toEqual({
      weekday: activeWeekHandlerService.activeWeek$[0],
      absences: [
        {
          personID: '1',
          reason: 'test',
          date: '2021-01-01',
          createdAt: jasmine.any(Date),
        },
      ],
    });
  });
});
