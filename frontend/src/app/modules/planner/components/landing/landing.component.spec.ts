import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LandingComponent } from './landing.component';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';

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
  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should set the department from the activated route', () => {
    expect(mockPlannerStateHandlerService.setActiveView).toHaveBeenCalledWith('department1', jasmine.any(Date));
  });

  it('should not set null department', () => {
    activatedRoute.queryParams = of({ department: null });
    component.ngOnInit();
    expect(
      mockPlannerStateHandlerService.setActiveView,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ).not.toHaveBeenCalledWith(null as any, jasmine.any(Date));
  });

  it('should turn department to lowercase', () => {
    activatedRoute.queryParams = of({ department: 'Department1' });
    component.ngOnInit();
    expect(mockPlannerStateHandlerService.setActiveView).toHaveBeenCalledWith('department1', jasmine.any(Date));
  });

  it('should not set undefined department', () => {
    activatedRoute.queryParams = of({ department: undefined });
    component.ngOnInit();
    expect(
      mockPlannerStateHandlerService.setActiveView,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ).not.toHaveBeenCalledWith(undefined as any, jasmine.any(Date));
  });
});
