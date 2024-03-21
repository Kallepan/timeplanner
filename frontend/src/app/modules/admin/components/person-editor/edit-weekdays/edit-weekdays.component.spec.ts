import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditWeekdaysComponent } from './edit-weekdays.component';
import { NotificationService } from '@app/core/services/notification.service';
import { PersonEditorDataContainerService } from '@app/modules/admin/services/person-editor-data-container.service';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { constants } from '@app/core/constants/constants';
import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatChipOptionHarness } from '@angular/material/chips/testing';
import { of } from 'rxjs';

describe('EditWeekdaysComponent', () => {
  let component: EditWeekdaysComponent;
  let fixture: ComponentFixture<EditWeekdaysComponent>;
  let loader: HarnessLoader;

  let mockPersonEditorDataContainerService: jasmine.SpyObj<PersonEditorDataContainerService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;

  beforeEach(async () => {
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);
    mockPersonEditorDataContainerService = jasmine.createSpyObj('PersonEditorDataContainerService', [], {
      weekdays$: constants.POSSIBLE_WEEKDAYS,
      activePerson$: null,
    });
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['addWeekdayToPerson', 'removeWeekdayFromPerson']);

    await TestBed.configureTestingModule({
      imports: [EditWeekdaysComponent],
      providers: [
        {
          provide: NotificationService,
          useValue: mockNotificationService,
        },
        {
          provide: PersonEditorDataContainerService,
          useValue: mockPersonEditorDataContainerService,
        },
        {
          provide: PersonAPIService,
          useValue: mockPersonAPIService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditWeekdaysComponent);
    loader = TestbedHarnessEnvironment.loader(fixture);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should handle null activePerson', async () => {
    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    await chips[0].toggle();
    expect(mockPersonAPIService.addWeekdayToPerson).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();

    await chips[0].toggle();
    expect(mockPersonAPIService.removeWeekdayFromPerson).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
  });
});

describe('EditWeekdaysComponent', () => {
  let component: EditWeekdaysComponent;
  let fixture: ComponentFixture<EditWeekdaysComponent>;
  let loader: HarnessLoader;

  let mockPersonEditorDataContainerService: jasmine.SpyObj<PersonEditorDataContainerService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;

  beforeEach(async () => {
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);
    mockPersonEditorDataContainerService = jasmine.createSpyObj('PersonEditorDataContainerService', [], {
      weekdays$: constants.POSSIBLE_WEEKDAYS,
      activePerson$: { id: '1', weekdays: [] },
    });
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['addWeekdayToPerson', 'removeWeekdayFromPerson']);

    await TestBed.configureTestingModule({
      imports: [EditWeekdaysComponent],
      providers: [
        {
          provide: NotificationService,
          useValue: mockNotificationService,
        },
        {
          provide: PersonEditorDataContainerService,
          useValue: mockPersonEditorDataContainerService,
        },
        {
          provide: PersonAPIService,
          useValue: mockPersonAPIService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditWeekdaysComponent);
    loader = TestbedHarnessEnvironment.loader(fixture);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render weekdays', async () => {
    expect(
      component.weekdays().map((weekday) => {
        return {
          id: weekday.id,
          name: weekday.name,
        };
      }),
    ).toEqual(constants.POSSIBLE_WEEKDAYS);

    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    expect(chips.length).toBe(constants.POSSIBLE_WEEKDAYS.length);
    chips.forEach(async (chip, index) => {
      expect(await chip.getText()).toBe(constants.POSSIBLE_WEEKDAYS[index].name);
    });
  });

  it('should select weekday', async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.addWeekdayToPerson.and.returnValue(of({} as any));
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.removeWeekdayFromPerson.and.returnValue(of({} as any));

    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    await chips[0].toggle();
    expect(mockPersonAPIService.addWeekdayToPerson).toHaveBeenCalledWith(constants.POSSIBLE_WEEKDAYS[0].id, '1');
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();

    await chips[0].toggle();
    expect(mockPersonAPIService.removeWeekdayFromPerson).toHaveBeenCalledWith(constants.POSSIBLE_WEEKDAYS[0].id, '1');
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should handle error', async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.addWeekdayToPerson.and.returnValue(of({} as any));
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.removeWeekdayFromPerson.and.returnValue(of({} as any));

    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    mockPersonAPIService.addWeekdayToPerson.and.throwError('error');
    await chips[0].toggle();
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();

    mockPersonAPIService.removeWeekdayFromPerson.and.throwError('error');
    await chips[0].toggle();
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();
  });
});
