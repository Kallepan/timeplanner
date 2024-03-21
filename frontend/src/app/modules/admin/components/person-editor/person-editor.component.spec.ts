import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonEditorComponent } from './person-editor.component';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { PersonEditorDataContainerService } from '../../services/person-editor-data-container.service';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { of } from 'rxjs';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { provideNoopAnimations } from '@angular/platform-browser/animations';

describe('PersonEditorComponent', () => {
  let component: PersonEditorComponent;
  let fixture: ComponentFixture<PersonEditorComponent>;

  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockWorkplaceAPIService: jasmine.SpyObj<WorkplaceAPIService>;
  let mockPersonDataContainerService: jasmine.SpyObj<PersonDataContainerService>;
  let personEditorDataContainerService: PersonEditorDataContainerService;

  const date = new Date();
  const mockPerson: PersonWithMetadata = {
    id: '1',
    first_name: 'John',
    last_name: 'Doe',
    email: 'test@example.com',
    active: true,
    working_hours: 8,
    workplaces: [],
    weekdays: [],
    deleted_at: null,
    created_at: date,
    updated_at: date,
  };
  const department = 'department';

  beforeEach(async () => {
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['getPerson']);
    mockPersonAPIService.getPerson.and.returnValue(of({ data: mockPerson, status: 200, message: 'OK' }));

    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [], {
      activeDepartment$: department,
    });

    mockWorkplaceAPIService = jasmine.createSpyObj('WorkplaceAPIService', ['getWorkplaces']);
    mockWorkplaceAPIService.getWorkplaces.and.returnValue(of({ data: [], status: 200, message: 'OK' }));

    mockPersonDataContainerService = jasmine.createSpyObj('PersonDataContainerService', ['activePerson$']);

    await TestBed.configureTestingModule({
      imports: [PersonEditorComponent],
      providers: [
        { provide: PersonAPIService, useValue: mockPersonAPIService },
        {
          provide: ActiveDepartmentHandlerService,
          useValue: mockActiveDepartmentHandlerService,
        },
        {
          provide: WorkplaceAPIService,
          useValue: mockWorkplaceAPIService,
        },
        {
          provide: PersonDataContainerService,
          useValue: mockPersonDataContainerService,
        },
        PersonEditorDataContainerService,
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonEditorComponent);
    personEditorDataContainerService = TestBed.inject(PersonEditorDataContainerService);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    expect(personEditorDataContainerService.activePerson$).toBeNull();
  });

  it('should set active person', () => {
    mockPersonAPIService.getPerson.and.returnValue(of({ data: mockPerson, status: 200, message: 'OK' }));
    component.setActivePerson('1');
    expect(mockPersonAPIService.getPerson).toHaveBeenCalledWith('1');
  });
});
