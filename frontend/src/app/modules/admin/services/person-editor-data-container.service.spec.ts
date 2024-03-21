import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonEditorDataContainerService } from './person-editor-data-container.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { of } from 'rxjs';
import { Component, inject } from '@angular/core';

@Component({
  template: `
    @if (service.workplaces$; as workplaces) {
      {{ workplaces.length }}
    }
  `,
  standalone: true,
})
class TestComponent {
  service = inject(PersonEditorDataContainerService);
}

describe('PersonEditorDataContainerService with Component', () => {
  let service: PersonEditorDataContainerService;
  let fixture: ComponentFixture<TestComponent>;
  const date = new Date();

  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockWorkplaceAPIService: jasmine.SpyObj<WorkplaceAPIService>;

  beforeEach(() => {
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [], {
      activeDepartment$: 'department',
    });
    mockWorkplaceAPIService = jasmine.createSpyObj('WorkplaceAPIService', ['getWorkplaces']);
    mockWorkplaceAPIService.getWorkplaces.and.returnValue(
      of({
        status: 200,
        data: [
          {
            id: '1',
            department_id: '1',
            name: 'workplace',
            created_at: date,
            updated_at: date,
            deleted_at: null,
          },
        ],
        message: 'OK',
      }),
    );

    TestBed.configureTestingModule({
      imports: [TestComponent],
      providers: [
        PersonEditorDataContainerService,
        {
          provide: ActiveDepartmentHandlerService,
          useValue: mockActiveDepartmentHandlerService,
        },
        {
          provide: WorkplaceAPIService,
          useValue: mockWorkplaceAPIService,
        },
      ],
    });
    service = TestBed.inject(PersonEditorDataContainerService);
    fixture = TestBed.createComponent(TestComponent);
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
    expect(mockWorkplaceAPIService.getWorkplaces).toHaveBeenCalledWith('department');
    expect(service.workplaces$).toEqual([
      {
        id: '1',
        department_id: '1',
        name: 'workplace',
        created_at: date,
        updated_at: date,
        deleted_at: null,
      },
    ]);
  });
});
