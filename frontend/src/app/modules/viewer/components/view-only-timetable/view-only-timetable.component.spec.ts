import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ViewOnlyTimetableComponent } from './view-only-timetable.component';

describe('ViewOnlyTimetableComponent', () => {
  let component: ViewOnlyTimetableComponent;
  let fixture: ComponentFixture<ViewOnlyTimetableComponent>;
  let mockActiveWeekdayService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockTimetableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;

  beforeEach(async () => {
    mockActiveWeekdayService = jasmine.createSpyObj('ActiveWeekHandlerService', [''], {
      activeWeek$: [],
    });

    mockTimetableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', [''], {
      timetableData$: [],
    });

    await TestBed.configureTestingModule({
      imports: [ViewOnlyTimetableComponent],
      providers: [
        {
          provide: ActiveWeekHandlerService,
          useValue: mockActiveWeekdayService,
        },
        {
          provide: TimetableDataContainerService,
          useValue: mockTimetableDataContainerService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(ViewOnlyTimetableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
