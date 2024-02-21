import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewerLandingPageComponent } from './viewer-landing-page.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ActivatedRoute } from '@angular/router';

describe('ViewerLandingPageComponent', () => {
  let component: ViewerLandingPageComponent;
  let fixture: ComponentFixture<ViewerLandingPageComponent>;

  // mocks
  let mockDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockActiveWeekHandlerService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockTimeTableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;

  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    mockDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['']);
    mockActiveWeekHandlerService = jasmine.createSpyObj('ActiveWeekHandlerService', ['']);
    mockTimeTableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', [''], {
      isLoading$: false,
    });

    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [''], {
      snapshot: {
        data: {
          title: 'Test',
          departmentId: 'departmentId',
        },
      },
    });

    await TestBed.configureTestingModule({
      providers: [
        { provide: ActiveDepartmentHandlerService, useValue: mockDepartmentHandlerService },
        { provide: ActiveWeekHandlerService, useValue: mockActiveWeekHandlerService },
        { provide: TimetableDataContainerService, useValue: mockTimeTableDataContainerService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      imports: [ViewerLandingPageComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ViewerLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
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
