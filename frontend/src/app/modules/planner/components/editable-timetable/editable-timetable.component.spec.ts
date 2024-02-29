import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditableTimetableComponent } from './editable-timetable.component';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { ActivatedRoute } from '@angular/router';

describe('EditableTimetableComponent', () => {
  let component: EditableTimetableComponent;
  let fixture: ComponentFixture<EditableTimetableComponent>;

  // mocks
  let mockPlannerStateHandlerService: jasmine.SpyObj<PlannerStateHandlerService>;
  let mockActiveWeekHandlerService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockTimetableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;
  let mockThemeHandlerService: jasmine.SpyObj<ThemeHandlerService>;

  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    mockPlannerStateHandlerService = jasmine.createSpyObj('PlannerStateHandlerService', ['assignPersonToTimeslot', 'unAssignPersonFromTimeslot']);
    mockActiveWeekHandlerService = jasmine.createSpyObj('ActiveWeekHandlerService', ['']);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['']);
    mockTimetableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', ['']);
    mockThemeHandlerService = jasmine.createSpyObj('ThemeHandlerService', ['']);

    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [''], {
      snapshot: {
        data: {
          title: 'Test',
          departmentId: 'departmentId',
        },
      },
    });

    await TestBed.configureTestingModule({
      imports: [EditableTimetableComponent],
      providers: [
        { provide: PlannerStateHandlerService, useValue: mockPlannerStateHandlerService },
        { provide: ActiveWeekHandlerService, useValue: mockActiveWeekHandlerService },
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
        { provide: TimetableDataContainerService, useValue: mockTimetableDataContainerService },
        { provide: ThemeHandlerService, useValue: mockThemeHandlerService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditableTimetableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call plannerStateHandlerService.assignPersonToTimeslot on personDroppedIntoTimeslotHandler', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const person = { id: 'id' } as any;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const timeslots = [{ id: 'id' } as any, { id: 'id' }] as any;
    component.personDroppedIntoTimeslotHandler(person, timeslots);
    expect(mockPlannerStateHandlerService.assignPersonToTimeslot).toHaveBeenCalledTimes(2);

    mockPlannerStateHandlerService.assignPersonToTimeslot.calls.reset();
    component.personDroppedIntoTimeslotHandler(person, []);
    expect(mockPlannerStateHandlerService.assignPersonToTimeslot).toHaveBeenCalledTimes(0);
  });

  it('should call plannerStateHandlerService.assignPersonToTimeslot on personAssignedToTimeslotEventHandler', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const person = { id: 'id' } as any;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const timeslot = { id: 'id' } as any;
    component.personAssignedToTimeslotEventHandler(person, timeslot);
    expect(mockPlannerStateHandlerService.assignPersonToTimeslot).toHaveBeenCalledTimes(1);
  });

  it('should call plannerStateHandlerService.unAssignPersonFromTimeslot on personUnassignedFromTimeslotEventHandler', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const person = { id: 'id' } as any;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const timeslot = { id: 'id' } as any;
    component.personUnassignedFromTimeslotEventHandler(person, timeslot);
    expect(mockPlannerStateHandlerService.unAssignPersonFromTimeslot).toHaveBeenCalledTimes(1);
  });

  it('should return slots from monday to friday', () => {
    const tests = [
      {
        slots: [
          { weekday: 1, name: 'Monday' },
          { weekday: 2, name: 'Tuesday' },
          { weekday: 3, name: 'Wednesday' },
          { weekday: 4, name: 'Thursday' },
          { weekday: 5, name: 'Friday' },
          { weekday: 6, name: 'Saturday' },
          { weekday: 7, name: 'Sunday' },
        ],
        expected: 5,
      },
      {
        slots: [
          { weekday: 1, name: 'Monday' },
          { weekday: 2, name: 'Tuesday' },
          { weekday: 3, name: 'Wednesday' },
          { weekday: 4, name: 'Thursday' },
          { weekday: 5, name: 'Friday' },
        ],
        expected: 5,
      },
      {
        slots: [
          { weekday: 6, name: 'Saturday' },
          { weekday: 7, name: 'Sunday' },
        ],
        expected: 0,
      },
      {
        slots: [],
        expected: 0,
      },
      {
        slots: [
          { weekday: 1, name: 'Monday' },
          { weekday: 2, name: 'Tuesday' },
          { weekday: 3, name: 'Wednesday' },
          { weekday: 4, name: 'Thursday' },
          { weekday: 5, name: 'Friday' },
          { weekday: 6, name: 'Saturday' },
          { weekday: 1, name: 'Monday' },
          { weekday: 2, name: 'Tuesday' },
          { weekday: 3, name: 'Wednesday' },
          { weekday: 4, name: 'Thursday' },
          { weekday: 5, name: 'Friday' },
          { weekday: 6, name: 'Saturday' },
        ],
        expected: 10,
      },
      {
        slots: [
          { weekday: 1, name: 'Monday' },
          { weekday: 2, name: 'Tuesday' },
          { weekday: 3, name: 'Wednesday' },
          { weekday: 4, name: 'Thursday' },
          { weekday: 5, name: 'Friday' },
          { weekday: 6, name: 'Saturday' },
          { weekday: 7, name: 'Sunday' },
          { weekday: 8, name: 'Invalid' },
          { weekday: 9, name: 'Invalid' },
          { weekday: 10, name: 'Invalid' },
          { weekday: 11, name: 'Invalid' },
          { weekday: 12, name: 'Invalid' },
        ],
        expected: 5,
      },
    ];

    tests.forEach((test) => {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const result = component.getSlotsFromMonToFri(test.slots as any[]);
      expect(result.length).toBe(test.expected);
    });
  });
});
