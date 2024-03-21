import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NotificationService } from '@app/core/services/notification.service';
import { PersonEditorDataContainerService } from '@app/modules/admin/services/person-editor-data-container.service';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatChipOptionHarness } from '@angular/material/chips/testing';
import { EditWorkplacesComponent } from './edit-workplaces.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { of } from 'rxjs';

describe('EditWorkplacesComponent', () => {
  let component: EditWorkplacesComponent;
  let fixture: ComponentFixture<EditWorkplacesComponent>;
  let loader: HarnessLoader;

  let mockPersonEditorDataContainerService: jasmine.SpyObj<PersonEditorDataContainerService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;

  beforeEach(async () => {
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);
    mockPersonEditorDataContainerService = jasmine.createSpyObj('PersonEditorDataContainerService', [], {
      workplaces$: [
        {
          id: '1',
          department_id: '1',
          name: 'workplace',
        },
        {
          id: '2',
          department_id: '2',
          name: 'workplace2',
        },
      ],
      activePerson$: null,
    });
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['addWorkplaceToPerson', 'removeWorkplaceFromPerson']);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [], {
      activeDepartment$: 'department',
    });

    await TestBed.configureTestingModule({
      imports: [EditWorkplacesComponent],
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
        {
          provide: ActiveDepartmentHandlerService,
          useValue: mockActiveDepartmentHandlerService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditWorkplacesComponent);
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
    expect(mockPersonAPIService.addWorkplaceToPerson).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();

    await chips[0].toggle();
    expect(mockPersonAPIService.removeWorkplaceFromPerson).not.toHaveBeenCalled();
    expect(mockNotificationService.infoMessage).not.toHaveBeenCalled();
  });
});

describe('EditWorkplacesComponent', () => {
  let component: EditWorkplacesComponent;
  let fixture: ComponentFixture<EditWorkplacesComponent>;
  let loader: HarnessLoader;

  let mockPersonEditorDataContainerService: jasmine.SpyObj<PersonEditorDataContainerService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;

  beforeEach(async () => {
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);
    mockPersonEditorDataContainerService = jasmine.createSpyObj('PersonEditorDataContainerService', [], {
      workplaces$: [
        {
          id: '1',
          department_id: '1',
          name: 'workplace',
        },
        {
          id: '2',
          department_id: '2',
          name: 'workplace2',
        },
      ],
      activePerson$: { id: '1', weekdays: [] },
    });
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['addWorkplaceToPerson', 'removeWorkplaceFromPerson']);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [], {
      activeDepartment$: 'department',
    });

    await TestBed.configureTestingModule({
      imports: [EditWorkplacesComponent],
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
        {
          provide: ActiveDepartmentHandlerService,
          useValue: mockActiveDepartmentHandlerService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditWorkplacesComponent);
    loader = TestbedHarnessEnvironment.loader(fixture);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render workplaces', async () => {
    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    expect(
      await Promise.all(
        chips.map(async (chip) => {
          return {
            name: await chip.getText(),
          };
        }),
      ),
    ).toEqual([{ name: 'workplace' }, { name: 'workplace2' }]);
  });

  it('should add workplace to person', async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.addWorkplaceToPerson.and.returnValue(of({} as any));
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.removeWorkplaceFromPerson.and.returnValue(of({} as any));
    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    await chips[0].toggle();
    expect(mockPersonAPIService.addWorkplaceToPerson).toHaveBeenCalledWith('department', '1', '1');
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should remove workplace from person', async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.addWorkplaceToPerson.and.returnValue(of({} as any));
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.removeWorkplaceFromPerson.and.returnValue(of({} as any));

    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    await chips[0].toggle();
    expect(mockPersonAPIService.addWorkplaceToPerson).toHaveBeenCalledWith('department', '1', '1');
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();

    await chips[0].toggle();
    expect(mockPersonAPIService.removeWorkplaceFromPerson).toHaveBeenCalledWith('department', '1', '1');
    expect(mockNotificationService.infoMessage).toHaveBeenCalled();
  });

  it('should handle error', async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.addWorkplaceToPerson.and.returnValue(of({} as any));
    mockPersonAPIService.addWorkplaceToPerson.and.throwError('error');

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mockPersonAPIService.removeWorkplaceFromPerson.and.returnValue(of({} as any));
    mockPersonAPIService.removeWorkplaceFromPerson.and.throwError('error');

    const chips = await loader.getAllHarnesses(MatChipOptionHarness);

    await chips[0].toggle();
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();

    await chips[0].toggle();
    expect(mockNotificationService.warnMessage).toHaveBeenCalled();
  });
});
