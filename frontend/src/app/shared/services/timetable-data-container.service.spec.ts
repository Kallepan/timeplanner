import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TimetableDataContainerService } from './timetable-data-container.service';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';
import { ActiveWeekHandlerService, Weekday } from './active-week-handler.service';
import { WorkdayAPIService } from './workday-api.service';
import { Component, WritableSignal, inject, signal } from '@angular/core';
import { mockWorkdays } from '@app/modules/viewer/tests/constants';
import { of } from 'rxjs';

describe('TimetableDataContainerService with Component', () => {
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockWorkdayAPIService: jasmine.SpyObj<WorkdayAPIService>;

  let fixture: ComponentFixture<TestComponent>;
  let service: TimetableDataContainerService;

  it('should not continue if activeDepartment is not set', () => {
    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays']);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [''], {
      activeDepartment$: null,
    });

    TestBed.configureTestingModule({
      imports: [TestComponent],
      providers: [
        TimetableDataContainerService,
        ActiveWeekHandlerService,
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
        { provide: WorkdayAPIService, useValue: mockWorkdayAPIService },
      ],
    });

    fixture = TestBed.createComponent(TestComponent);
    TestBed.inject(ActiveWeekHandlerService);
    service = TestBed.inject(TimetableDataContainerService);
    fixture.detectChanges();

    expect(service.workplaces$).toEqual([]);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);

    expect(mockWorkdayAPIService.getWorkdays).not.toHaveBeenCalled();
  });
});
@Component({
  selector: 'app-timetable-data-container',
  template: `
    @if (!!getWorkplaces()) {
      <div></div>
    }
    @if (!!listOfPersonsAssignedToTheWholeWeek()) {
      <div></div>
    }
  `,
  standalone: true,
})
class TestComponent {
  timetableDataContainerService = inject(TimetableDataContainerService);

  // we dont need these wrappers, but who cares
  getWorkplaces() {
    this.timetableDataContainerService.workplaces$;
  }
  listOfPersonsAssignedToTheWholeWeek() {
    this.timetableDataContainerService.listOfPersonsAssignedToTheWholeWeek();
  }
}

describe('TimetableDataContainerService with Component', () => {
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockWorkdayAPIService: jasmine.SpyObj<WorkdayAPIService>;

  let activeWeekHandlerService: ActiveWeekHandlerService;
  let fixture: ComponentFixture<TestComponent>;
  let service: TimetableDataContainerService;

  beforeEach(() => {
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [''], {
      activeDepartment$: 'test',
    });

    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays']);

    TestBed.configureTestingModule({
      imports: [TestComponent],
      providers: [
        TimetableDataContainerService,
        ActiveWeekHandlerService,
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
        { provide: WorkdayAPIService, useValue: mockWorkdayAPIService },
      ],
    });
    fixture = TestBed.createComponent(TestComponent);
    activeWeekHandlerService = TestBed.inject(ActiveWeekHandlerService);
    service = TestBed.inject(TimetableDataContainerService);
    fixture.detectChanges();
  });

  it('should handle emptyy workplaces$', () => {
    expect(service.workplaces$).toEqual([]);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should set workplaces$', () => {
    expect(service.workplaces$).toEqual([]);
    mockWorkdayAPIService.getWorkdays.and.returnValue(of({ data: mockWorkdays, message: 'success', status: 200 }));
    activeWeekHandlerService.activeWeekByDate = new Date('2024-01-01');

    fixture.detectChanges();

    expect(mockWorkdayAPIService.getWorkdays).toHaveBeenCalled();
    expect(service.workplaces$).not.toEqual([]);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
  });

  it('should have expected workplaces$ output', () => {
    expect(service.workplaces$).toEqual([]);
    mockWorkdayAPIService.getWorkdays.and.returnValue(of({ data: mockWorkdays, message: 'success', status: 200 }));
    activeWeekHandlerService.activeWeekByDate = new Date();

    fixture.detectChanges();

    expect(mockWorkdayAPIService.getWorkdays).toHaveBeenCalled();
    const parsedData = JSON.stringify(service.workplaces$);
    expect(parsedData).toContain('gridRow');
    expect(parsedData).toContain('colorForLightMode');
    expect(parsedData).toContain('colorForDarkMode');
    expect(service.workplaces$).not.toEqual([]);
    expect(service.workplaces$.length).toBe(1);
    expect(service.workplaces$[0].timeslotGroups.length).toBe(1);
    expect(service.workplaces$[0].timeslotGroups[0].workdayTimeslots.length).toBe(14);
  });

  it('should add person to the list of persons assigned to the "whole" (MON to FRI) week', () => {
    mockWorkdayAPIService.getWorkdays.and.returnValue(of({ data: mockWorkdays, message: 'success', status: 200 }));
    activeWeekHandlerService.activeWeekByDate = new Date();

    fixture.detectChanges();

    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);

    // populate the list of persons assigned to the whole week
    service.addPersonWithWeekday('1', 1);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
    service.addPersonWithWeekday('1', 2);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
    service.addPersonWithWeekday('1', 3);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
    service.addPersonWithWeekday('1', 4);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);

    // should not add person to the list if they are assigned to weekend
    service.addPersonWithWeekday('1', 6);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
    service.addPersonWithWeekday('1', 7);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);

    // should add person to the list if weekdays 1 to 5 are present
    service.addPersonWithWeekday('1', 5);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual(['1']);
  });

  it('should remove person from the list of persons assigned to the "whole" (MON to FRI) week', () => {
    mockWorkdayAPIService.getWorkdays.and.returnValue(of({ data: mockWorkdays, message: 'success', status: 200 }));
    activeWeekHandlerService.activeWeekByDate = new Date();

    fixture.detectChanges();

    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);

    // populate the list of persons assigned to the whole week
    service.addPersonWithWeekday('1', 1);
    service.addPersonWithWeekday('1', 2);
    service.addPersonWithWeekday('1', 3);
    service.addPersonWithWeekday('1', 4);
    service.addPersonWithWeekday('1', 5);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual(['1']);

    // remove person from the list
    service.removePersonWithWeekday('1', 5);
    expect(service.listOfPersonsAssignedToTheWholeWeek()).toEqual([]);
  });
});

describe('TimetableDataContainerService', () => {
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockActiveWeekHandlerService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockWorkdayAPIService: jasmine.SpyObj<WorkdayAPIService>;

  let service: TimetableDataContainerService;

  let mockActiveWeekSignal: WritableSignal<Weekday[]>;
  beforeEach(() => {
    mockActiveWeekSignal = signal([]);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['']);
    mockActiveWeekHandlerService = jasmine.createSpyObj('ActiveWeekHandlerService', [''], {
      activeWeek$: mockActiveWeekSignal,
    });
    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays']);

    TestBed.configureTestingModule({
      providers: [
        TimetableDataContainerService,
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
        { provide: ActiveWeekHandlerService, useValue: mockActiveWeekHandlerService },
        { provide: WorkdayAPIService, useValue: mockWorkdayAPIService },
      ],
    });
    service = TestBed.inject(TimetableDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should have a default value for colorizeMissing$', () => {
    expect(service.colorizeMissing$).toBeTrue();
  });

  it('should have a default value for isLoading$', () => {
    expect(service.isLoading$).toBeTrue();
  });

  it('should have a default value for displayComments$', () => {
    expect(service.displayComments$).toBeTrue();
  });

  it('should have a default value for displayTimes$', () => {
    expect(service.displayTimes$).toBeFalse();
  });

  it('should have a default value for colorize$', () => {
    expect(service.colorize$).toBeTrue();
  });

  it('should have a default value for workplaces$', () => {
    expect(service.workplaces$).toEqual([]);
  });

  it('should set colorizeMissing$', () => {
    expect(service.colorizeMissing$).toBeTrue();
    service.colorizeMissing = false;
    expect(service.colorizeMissing$).toBeFalse();
  });

  it('should set displayComments$', () => {
    expect(service.displayComments$).toBeTrue();
    service.displayComments = false;
    expect(service.displayComments$).toBeFalse();
  });

  it('should set displayTimes$', () => {
    expect(service.displayTimes$).toBeFalse();
    service.displayTimes = true;
    expect(service.displayTimes$).toBeTrue();
  });

  it('should set colorize$', () => {
    expect(service.colorize$).toBeTrue();
    service.colorize = false;
    expect(service.colorize$).toBeFalse();
  });
});
