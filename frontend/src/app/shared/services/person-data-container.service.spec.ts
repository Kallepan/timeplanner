import { ComponentFixture, TestBed } from '@angular/core/testing';

import { of } from 'rxjs';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';
import { PersonAPIService } from './person-api.service';
import { PersonDataContainerService } from './person-data-container.service';
import { ActivatedRoute } from '@angular/router';
import { Component, inject } from '@angular/core';

describe('PersonDataContainerService', () => {
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let service: PersonDataContainerService;

  beforeEach(() => {
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['getPersons']);
    mockPersonAPIService.getPersons.and.returnValue(of({ data: [], message: 'test', status: 200 }));
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      queryParams: of({
        department: 'test2',
      }),
    });

    TestBed.configureTestingModule({
      providers: [
        PersonDataContainerService,
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: PersonAPIService, useValue: mockPersonAPIService },
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
      ],
    });
    service = TestBed.inject(PersonDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should have a default persons', () => {
    expect(service.persons$).toEqual([]);
  });
});

// Crate a mock component and subscribe to the persons$ signal
@Component({
  template: `{{ persons$ }}`,
})
class PersonDataContainerServiceTestComponent {
  private personDataContainerService = inject(PersonDataContainerService);
  persons$ = this.personDataContainerService.persons$;
}

describe('PersonDataContainerServiceTestComponent', () => {
  let fixture: ComponentFixture<PersonDataContainerServiceTestComponent>;
  let mockPersonAPIService: jasmine.SpyObj<PersonAPIService>;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  let activeDepartmentHandlerService: ActiveDepartmentHandlerService;
  let personDataContainerService: PersonDataContainerService;

  beforeEach(async () => {
    mockPersonAPIService = jasmine.createSpyObj('PersonAPIService', ['getPersons']);
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      queryParams: of({
        department: 'test2',
      }),
    });
    mockPersonAPIService.getPersons.and.returnValue(of({ data: [], message: 'test2', status: 200 }));

    await TestBed.configureTestingModule({
      declarations: [PersonDataContainerServiceTestComponent],
      providers: [PersonDataContainerService, { provide: ActivatedRoute, useValue: mockActivatedRoute }, { provide: PersonAPIService, useValue: mockPersonAPIService }, ActiveDepartmentHandlerService],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonDataContainerServiceTestComponent);

    activeDepartmentHandlerService = TestBed.inject(ActiveDepartmentHandlerService);
    personDataContainerService = TestBed.inject(PersonDataContainerService);
  });

  it('should update persons when active department changes', () => {
    expect(personDataContainerService.persons$).toEqual([]);
    const department = 'test';
    const persons = [
      {
        id: '1',
        first_name: 'test',
        last_name: 'test',
        email: 'test@example.com',
        active: true,
        working_hours: 8,
        workplaces: [],
        departments: [],
        weekdays: [],
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
    ];

    mockPersonAPIService.getPersons.and.returnValue(of({ data: persons, message: 'test', status: 200 }));
    activeDepartmentHandlerService.activeDepartment = department;

    fixture.detectChanges();

    expect(mockPersonAPIService.getPersons).toHaveBeenCalledWith(department);
    expect(personDataContainerService.persons$).toEqual(persons);
  });
});
