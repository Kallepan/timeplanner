import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { LandingComponent } from './landing.component';

describe('LandingComponent', () => {
  let component: LandingComponent;
  let fixture: ComponentFixture<LandingComponent>;
  let mockPlannerStateHandlerService: jasmine.SpyObj<PlannerStateHandlerService>;
  let mockTimetableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;

  beforeEach(async () => {
    mockPlannerStateHandlerService = jasmine.createSpyObj('PlannerStateHandlerService', ['setActiveView']);
    mockTimetableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', [''], {
      colorize: true,
      displaytimes: true,
      displayTimes$: true,
      colorize$: true,
    });

    await TestBed.configureTestingModule({
      imports: [LandingComponent],
      providers: [
        {
          provide: PlannerStateHandlerService,
          useValue: mockPlannerStateHandlerService,
        },
        {
          provide: TimetableDataContainerService,
          useValue: mockTimetableDataContainerService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(LandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
