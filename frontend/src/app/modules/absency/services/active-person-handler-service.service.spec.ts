import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ActivePersonHandlerServiceService } from './active-person-handler-service.service';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { Component, inject } from '@angular/core';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { of, throwError } from 'rxjs';
import CalendarDayEventObject from 'js-year-calendar/dist/interfaces/CalendarDayEventObject';
import CalendarDataSourceElement from 'js-year-calendar/dist/interfaces/CalendarDataSourceElement';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { NotificationService } from '@app/core/services/notification.service';
import { formatDateToDateString } from '@app/shared/functions/format-date-to-string.function';

const mockPerson: PersonWithMetadata = {
  created_at: new Date(),
  updated_at: new Date(),
  deleted_at: null,
  id: '123',
  first_name: 'John',
  last_name: 'Doe',
  email: 'doe@example.com',
  active: true,
  working_hours: 40,
  workplaces: [],
  departments: [],
  weekdays: [],
};

describe('ActivePersonHandlerServiceService', () => {
  let service: ActivePersonHandlerServiceService;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;

  beforeEach(() => {
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['getAbsencyForPersonInRange']);

    TestBed.configureTestingModule({
      providers: [
        ActivePersonHandlerServiceService,
        {
          provide: PersonAPIService,
          useValue: mockPersonAPIService,
        },
      ],
    });
    service = TestBed.inject(ActivePersonHandlerServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should have default active person of null', () => {
    expect(service.activePerson$).toBeNull();
  });

  it('should set active person', () => {
    service.activePerson = mockPerson;

    expect(service.activePerson$).toBe(mockPerson);
  });

  it('should have default active year of current year', () => {
    expect(service.activeYear$).toBe(new Date().getFullYear());
  });

  it('should set active year', () => {
    const year = 2021;
    service.activeYear = year;

    expect(service.activeYear$).toBe(year);
  });

  it('should have default absences of empty array', () => {
    expect(service.absences$).toEqual([]);
  });
});

@Component({
  template: ` @for (absence of service.absences$; track absence) {
      {{ absence.name }}
      <!-- name is a property of absence renamed from reason -->
    } @empty {
      none
    }
    <br />{{ service.activeYear$ }}`,
})
class TestComponent {
  service = inject(ActivePersonHandlerServiceService);
}

describe('ActivePersonHandlerServiceService in TestComponent', () => {
  let service: ActivePersonHandlerServiceService;
  let fixture: ComponentFixture<TestComponent>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;

  beforeEach(() => {
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['getAbsencyForPersonInRange', 'addAbsencyToPerson', 'removeAbsencyFromPerson']);
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['warnMessage', 'infoMessage']);
    TestBed.configureTestingModule({
      declarations: [TestComponent],
      providers: [
        ActivePersonHandlerServiceService,
        {
          provide: PersonAPIService,
          useValue: mockPersonAPIService,
        },
        {
          provide: NotificationService,
          useValue: mockNotificationService,
        },

        provideNoopAnimations(),
      ],
    });
    service = TestBed.inject(ActivePersonHandlerServiceService);
    fixture = TestBed.createComponent(TestComponent);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
    expect(fixture).toBeTruthy();
  });

  it('should have default active person of null', () => {
    expect(service.activePerson$).toBeNull();
  });

  it('should call getAbsencyForPersonInRange when active year changes', () => {
    mockPersonAPIService.getAbsencyForPersonInRange.and.returnValue(of({ data: [], message: 'success', status: 200 }));
    const year = 2021;
    service.activePerson = mockPerson;
    service.activeYear = year;

    fixture.detectChanges();
    expect(fixture.nativeElement.textContent).toContain(year);
    expect(fixture.nativeElement.textContent).toContain('none');
    expect(mockPersonAPIService.getAbsencyForPersonInRange).toHaveBeenCalledWith(mockPerson.id, `${year}-01-01`, `${year}-12-31`);

    mockPersonAPIService.getAbsencyForPersonInRange.calls.reset();
    mockPersonAPIService.getAbsencyForPersonInRange.and.returnValue(
      of({
        data: [
          {
            person_id: mockPerson.id,
            reason: 'sick',
            date: '2021-01-01',
            created_at: new Date(),
          },
        ],
        message: 'success',
        status: 200,
      }),
    );
    service.activeYear = 2022;

    fixture.detectChanges();

    expect(mockPersonAPIService.getAbsencyForPersonInRange).toHaveBeenCalledWith(mockPerson.id, `2022-01-01`, `2022-12-31`);
    expect(fixture.nativeElement.textContent).toContain('2022');
    expect(fixture.nativeElement.textContent).toContain('sick');
  });

  it('should not call getAbsencyForPersonInRange when active person is null', () => {
    service.activeYear = 2021;

    fixture.detectChanges();
    expect(fixture.nativeElement.textContent).toContain('none');
    expect(mockPersonAPIService.getAbsencyForPersonInRange).not.toHaveBeenCalled();
  });

  it('should handle Remove Absency with null return value', () => {
    const currentDate = new Date();
    const mockCalendarDayEventObject: CalendarDayEventObject<CalendarDataSourceElement> = {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      element: {} as any,
      events: [],
      date: currentDate,
    };
    service.activePerson = mockPerson;
    service.activeYear = currentDate.getFullYear();
    // mock API response
    mockPersonAPIService.removeAbsencyFromPerson.and.returnValue(of({ data: null, message: 'success', status: 200 }));
    mockPersonAPIService.getAbsencyForPersonInRange.and.returnValue(
      of({
        data: [
          {
            person_id: mockPerson.id,
            reason: 'sick',
            date: formatDateToDateString(currentDate),
            created_at: new Date(),
          },
        ],
        message: 'success',
        status: 200,
      }),
    );
    const dialogSpy = spyOn(service.dialog, 'open');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    dialogSpy.and.returnValue({
      afterClosed: () => of(true),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    service.removeAbsency(mockCalendarDayEventObject);
    fixture.detectChanges();

    expect(mockPersonAPIService.removeAbsencyFromPerson).toHaveBeenCalledWith(mockPerson.id, formatDateToDateString(currentDate));
    expect(mockPersonAPIService.getAbsencyForPersonInRange).toHaveBeenCalledTimes(2);
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should handle Add Absency', () => {
    const currentDate = new Date();
    service.activePerson = mockPerson;
    service.activeYear = currentDate.getFullYear();
    // mock API response
    mockPersonAPIService.addAbsencyToPerson.and.returnValue(of({ data: null, message: 'success', status: 200 }));
    mockPersonAPIService.getAbsencyForPersonInRange.and.returnValue(
      of({
        data: [
          {
            person_id: mockPerson.id,
            reason: 'sick',
            date: formatDateToDateString(currentDate),
            created_at: new Date(),
          },
        ],
        message: 'success',
        status: 200,
      }),
    );
    const dialogSpy = spyOn(service.dialog, 'open');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    dialogSpy.and.returnValue({
      afterClosed: () =>
        of({
          endDate: currentDate,
          reason: 'sick',
        }),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    const mockCalendarDayEventObject: CalendarDayEventObject<CalendarDataSourceElement> = {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      element: {} as any,
      events: [],
      date: currentDate,
    };

    service.addAbsency(mockCalendarDayEventObject);
    fixture.detectChanges();

    expect(mockPersonAPIService.addAbsencyToPerson).toHaveBeenCalledWith(mockPerson.id, formatDateToDateString(currentDate), 'sick');
    expect(mockPersonAPIService.getAbsencyForPersonInRange).toHaveBeenCalledTimes(2);
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should handle error during remove absency', () => {
    const currentDate = new Date();
    service.activePerson = mockPerson;
    service.activeYear = currentDate.getFullYear();
    // mock API response
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.removeAbsencyFromPerson.and.returnValue(throwError(() => new Error({ status: 500 } as any)));
    mockPersonAPIService.getAbsencyForPersonInRange.and.returnValue(of({ data: [], message: 'success', status: 200 }));
    const dialogSpy = spyOn(service.dialog, 'open');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    dialogSpy.and.returnValue({
      afterClosed: () => of(true),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    const mockCalendarDayEventObject: CalendarDayEventObject<CalendarDataSourceElement> = {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      element: {} as any,
      events: [],
      date: currentDate,
    };

    service.removeAbsency(mockCalendarDayEventObject);
    fixture.detectChanges();

    expect(mockPersonAPIService.removeAbsencyFromPerson).toHaveBeenCalledWith(mockPerson.id, formatDateToDateString(currentDate));
    expect(mockPersonAPIService.getAbsencyForPersonInRange).toHaveBeenCalledTimes(1); // once during initialization
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();
  });

  it('should handle error during add absency', () => {
    const currentDate = new Date();
    service.activePerson = mockPerson;
    service.activeYear = currentDate.getFullYear();

    // mock API response
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.addAbsencyToPerson.and.returnValue(throwError(() => new Error({ status: 500 } as any)));
    mockPersonAPIService.getAbsencyForPersonInRange.and.returnValue(
      of({
        data: [
          {
            person_id: mockPerson.id,
            reason: 'sick',
            date: formatDateToDateString(currentDate),
            created_at: new Date(),
          },
        ],
        message: 'success',
        status: 200,
      }),
    );
    const dialogSpy = spyOn(service.dialog, 'open');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    dialogSpy.and.returnValue({
      afterClosed: () =>
        of({
          endDate: currentDate,
          reason: 'sick',
        }),
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);

    const mockCalendarDayEventObject: CalendarDayEventObject<CalendarDataSourceElement> = {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      element: {} as any,
      events: [],
      date: currentDate,
    };

    service.addAbsency(mockCalendarDayEventObject);
    fixture.detectChanges();

    expect(mockPersonAPIService.addAbsencyToPerson).toHaveBeenCalledWith(mockPerson.id, formatDateToDateString(currentDate), 'sick');
    expect(mockPersonAPIService.getAbsencyForPersonInRange).toHaveBeenCalledTimes(1); // once during initialization
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();
  });
});
