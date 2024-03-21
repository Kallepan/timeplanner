import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonEditorLandingPageComponent } from './person-editor-landing-page.component';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { of } from 'rxjs';

describe('PersonEditorLandingPageComponent', () => {
  let component: PersonEditorLandingPageComponent;
  let fixture: ComponentFixture<PersonEditorLandingPageComponent>;

  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;

  beforeEach(async () => {
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['getDepartments']);
    mockDepartmentAPIService.getDepartments.and.returnValue(of({ data: [], message: 'Success', status: 200 }));

    await TestBed.configureTestingModule({
      imports: [PersonEditorLandingPageComponent],
      providers: [{ provide: DepartmentAPIService, useValue: mockDepartmentAPIService }],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonEditorLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
