import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewerLandingPageComponent } from './viewer-landing-page.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ActivatedRoute } from '@angular/router';
import { MatBottomSheet, MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { AbsencyPanelComponent } from '@app/modules/planner/components/absency-panel/absency-panel.component';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';

describe('ViewerLandingPageComponent', () => {
  let component: ViewerLandingPageComponent;
  let fixture: ComponentFixture<ViewerLandingPageComponent>;

  // mocks
  let mockDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockActiveWeekHandlerService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockTimeTableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;

  let mockBottomSheet: jasmine.SpyObj<MatBottomSheet>;

  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    mockDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['']);
    mockActiveWeekHandlerService = jasmine.createSpyObj('ActiveWeekHandlerService', ['']);
    mockTimeTableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', [''], {
      isLoading$: false,
    });
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['']);

    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [''], {
      snapshot: {
        data: {
          title: 'Test',
          departmentId: 'departmentId',
        },
      },
    });

    mockBottomSheet = jasmine.createSpyObj('MatBottomSheet', ['open']);

    await TestBed.configureTestingModule({
      providers: [
        { provide: ActiveDepartmentHandlerService, useValue: mockDepartmentHandlerService },
        { provide: ActiveWeekHandlerService, useValue: mockActiveWeekHandlerService },
        { provide: TimetableDataContainerService, useValue: mockTimeTableDataContainerService },
        { provide: DepartmentAPIService, useValue: mockDepartmentAPIService },
        { provide: MatBottomSheet, useValue: mockBottomSheet },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        provideNoopAnimations(),
      ],
      imports: [ViewerLandingPageComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ViewerLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();

    expect(mockActiveWeekHandlerService.activeWeekByDate.toDateString()).toBe(new Date().toDateString());
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
});
