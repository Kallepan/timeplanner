import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LandingPageComponent } from './landing-page.component';
import { ViewerStateHandlerService } from '../../services/viewer-state-handler.service';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';

describe('LandingPageComponent', () => {
  let component: LandingPageComponent;
  let fixture: ComponentFixture<LandingPageComponent>;
  let mockViewerStateHandlerService: jasmine.SpyObj<ViewerStateHandlerService>;
  let activatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    activatedRoute = jasmine.createSpyObj('ActivatedRoute', ['data'], {
      queryParams: of({ department: 'department1' }),
    });
    mockViewerStateHandlerService = jasmine.createSpyObj('ViewerStateHandlerService', ['setActiveView']);

    await TestBed.configureTestingModule({
      providers: [
        {
          provide: ViewerStateHandlerService,
          useValue: mockViewerStateHandlerService,
        },
        {
          provide: ActivatedRoute,
          useValue: activatedRoute,
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

  it('should set the department from the activated route', () => {
    expect(mockViewerStateHandlerService.setActiveView).toHaveBeenCalledWith('department1', jasmine.any(Date));
  });

  it('should not set null department', () => {
    activatedRoute.queryParams = of({ department: null });
    component.ngOnInit();
    expect(
      mockViewerStateHandlerService.setActiveView,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ).not.toHaveBeenCalledWith(null as any, jasmine.any(Date));
  });

  it('should turn department to lowercase', () => {
    activatedRoute.queryParams = of({ department: 'Department1' });
    component.ngOnInit();
    expect(mockViewerStateHandlerService.setActiveView).toHaveBeenCalledWith('department1', jasmine.any(Date));
  });

  it('should not set undefined department', () => {
    activatedRoute.queryParams = of({ department: undefined });
    component.ngOnInit();
    expect(
      mockViewerStateHandlerService.setActiveView,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ).not.toHaveBeenCalledWith(undefined as any, jasmine.any(Date));
  });
});
