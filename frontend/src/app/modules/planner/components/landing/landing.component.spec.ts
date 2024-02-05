import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { LandingComponent } from './landing.component';

describe('LandingComponent', () => {
  let component: LandingComponent;
  let fixture: ComponentFixture<LandingComponent>;
  let mockPlannerStateHandlerService: jasmine.SpyObj<PlannerStateHandlerService>;
  let activatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    activatedRoute = jasmine.createSpyObj('ActivatedRoute', ['data'], {
      queryParams: of({ department: 'department1' }),
    });
    mockPlannerStateHandlerService = jasmine.createSpyObj('PlannerStateHandlerService', ['setActiveView']);

    await TestBed.configureTestingModule({
      imports: [LandingComponent],
      providers: [
        {
          provide: PlannerStateHandlerService,
          useValue: mockPlannerStateHandlerService,
        },
        {
          provide: ActivatedRoute,
          useValue: activatedRoute,
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
