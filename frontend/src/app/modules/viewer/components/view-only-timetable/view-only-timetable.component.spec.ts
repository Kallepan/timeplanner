import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ViewOnlyTimetableComponent } from './view-only-timetable.component';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';

describe('ViewOnlyTimetableComponent', () => {
  let component: ViewOnlyTimetableComponent;
  let fixture: ComponentFixture<ViewOnlyTimetableComponent>;
  let mockActiveWeekdayService: jasmine.SpyObj<ActiveWeekHandlerService>;
  let mockTimetableDataContainerService: jasmine.SpyObj<TimetableDataContainerService>;
  let mockThemeHandlerService: jasmine.SpyObj<ThemeHandlerService>;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    mockActiveWeekdayService = jasmine.createSpyObj('ActiveWeekHandlerService', [''], {
      activeWeek$: [],
    });

    mockTimetableDataContainerService = jasmine.createSpyObj('TimetableDataContainerService', [''], {
      timetableData$: [],
    });

    mockThemeHandlerService = jasmine.createSpyObj('ThemeHandlerService', ['toggleTheme'], {
      isDark$: true,
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
        {
          provide: ThemeHandlerService,
          useValue: mockThemeHandlerService,
        },
        provideHttpClient(),
        provideHttpClientTesting(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(ViewOnlyTimetableComponent);
    httpMock = TestBed.inject(HttpTestingController);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch css file and open print window', () => {
    expect(component.print).toBeDefined();

    const element = document.createElement('div');
    element.innerHTML = 'test';

    component.print(element);

    const req = httpMock.expectOne('assets/print-styles.css');
    expect(req.request.method).toBe('GET');
    req.flush('test');

    expect(mockThemeHandlerService.toggleTheme).toHaveBeenCalledTimes(2);
  });
});
