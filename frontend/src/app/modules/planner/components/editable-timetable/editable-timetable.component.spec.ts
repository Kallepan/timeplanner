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
    const slots = [
      { id: 'id', weekday: 'MON' },
      { id: 'id', weekday: 'TUE' },
      { id: 'id', weekday: 'WED' },
      { id: 'id', weekday: 'THU' },
      { id: 'id', weekday: 'FRI' },
      { id: 'id', weekday: 'SAT' },
      { id: 'id', weekday: 'SUN' },
    ];

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const result = component.getSlotsFromMonToFri(slots as any[]);
    expect(result.length).toBe(5);
  });
});
