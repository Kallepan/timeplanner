import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { LandingPageComponent } from './landing-page.component';

describe('LandingPageComponent', () => {
  let component: LandingPageComponent;
  let fixture: ComponentFixture<LandingPageComponent>;
  let mockTimetableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;
  let mockActiveWeekHandlerService: jasmine.SpyObj<ActiveWeekHandlerService>;

  beforeEach(async () => {
    mockTimetableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', ['getTimetableData']);
    mockActiveWeekHandlerService = jasmine.createSpyObj('ActiveWeekHandlerService', ['shiftWeek']);

    await TestBed.configureTestingModule({
      providers: [
        {
          provide: TimetableDataContainerService,
          useValue: mockTimetableDataContainerService,
        },
        {
          provide: ActiveWeekHandlerService,
          useValue: mockActiveWeekHandlerService,
        },
      ],
      imports: [LandingPageComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(LandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
