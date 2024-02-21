import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PlannerLandingPageComponent } from './planner-landing-page.component';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { ActivatedRoute } from '@angular/router';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { MatBottomSheet, MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { AbsencyPanelComponent } from '../absency-panel/absency-panel.component';

describe('PlannerLandingPageComponent', () => {
  let component: PlannerLandingPageComponent;
  let fixture: ComponentFixture<PlannerLandingPageComponent>;

  // mocks
  let mockTimeTableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;
  let mockActiveWeekHandlerService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockPlannerStateHandlerService: jasmine.SpyObj<PlannerStateHandlerService>;

  let mockBottomSheet: jasmine.SpyObj<MatBottomSheet>;

  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    mockTimeTableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', [''], {
      isLoading$: false,
    });
    mockActiveWeekHandlerService = jasmine.createSpyObj('ActiveWeekHandlerService', ['']);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['']);
    mockPlannerStateHandlerService = jasmine.createSpyObj('PlannerStateHandlerService', ['']);

    mockBottomSheet = jasmine.createSpyObj('MatBottomSheet', ['open']);

    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [''], {
      snapshot: {
        data: {
          title: 'Test',
          departmentId: 'departmentId',
        },
      },
    });

    await TestBed.configureTestingModule({
      imports: [PlannerLandingPageComponent],
      providers: [
        { provide: TimetableDataContainerService, useValue: mockTimeTableDataContainerService },
        { provide: ActiveWeekHandlerService, useValue: mockActiveWeekHandlerService },
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
        { provide: PlannerStateHandlerService, useValue: mockPlannerStateHandlerService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: MatBottomSheet, useValue: mockBottomSheet },
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(PlannerLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should toggle absency panel', () => {
    component.toggleAbsencyPanel();

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    expect(mockBottomSheet.open).toHaveBeenCalledWith(AbsencyPanelComponent as any, { hasBackdrop: false, closeOnNavigation: true });
  });

  it('should dismiss absency panel', () => {
    // return a dummy object
    mockBottomSheet.open.and.returnValue({ dismiss: () => {} } as MatBottomSheetRef);

    component.toggleAbsencyPanel();
    fixture.detectChanges();
    expect(component._bottomSheetRef).not.toBeNull();

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const bottomSheetRefSpy = spyOn(component._bottomSheetRef as any, 'dismiss');

    component.toggleAbsencyPanel();

    expect(bottomSheetRefSpy).toHaveBeenCalledTimes(1);
    expect(mockBottomSheet.open).toHaveBeenCalledTimes(1);
  });

  it('should have an overlay', () => {
    const loadingStatusSpy = spyOn(component, 'getLoadingStatus').and.callThrough();
    loadingStatusSpy.and.returnValue(false);

    const overlay = fixture.nativeElement.querySelector('#overlay');

    expect(overlay).toBeTruthy();
    expect(overlay.textContent).toBe(''); // empty overlay

    // check if styles are applied
    expect(overlay.style.display).toBe('none');

    // change loading status
    loadingStatusSpy.and.returnValue(true);
    fixture.detectChanges();

    expect(overlay.style.display).toBe('block');
  });
});
