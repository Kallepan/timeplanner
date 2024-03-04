import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SchemaEditDialogComponent, SchemaEditDialogData } from './schema-edit-dialog.component';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { TimeslotAPIService } from '@app/shared/services/timeslot-api.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { provideNoopAnimations } from '@angular/platform-browser/animations';

describe('SchemaEditDialogComponent with DepartmentServiceValidator', () => {
  let component: SchemaEditDialogComponent<DepartmentAPIService>;
  let fixture: ComponentFixture<SchemaEditDialogComponent<DepartmentAPIService>>;
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;

  beforeEach(async () => {
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['create', 'update', 'delete']);

    const mockMatDialogData: SchemaEditDialogData<DepartmentAPIService> = {
      id: '',
      name: '',
      idIsEditable: false,
      serviceForValidation: mockDepartmentAPIService,
    };
    await TestBed.configureTestingModule({
      imports: [SchemaEditDialogComponent],
      providers: [{ provide: MAT_DIALOG_DATA, useValue: mockMatDialogData }, provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(SchemaEditDialogComponent<DepartmentAPIService>);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should disable id input if idIsEditable is false', () => {
    expect(component.formGroup.controls.id.disabled).toBeTrue();
  });
});

describe('SchemaEditDialogComponent with TimeslotServiceValidator', () => {
  let component: SchemaEditDialogComponent<TimeslotAPIService>;
  let fixture: ComponentFixture<SchemaEditDialogComponent<TimeslotAPIService>>;
  let mockTimeslotAPIService: jasmine.SpyObj<TimeslotAPIService>;

  beforeEach(async () => {
    mockTimeslotAPIService = jasmine.createSpyObj('TimeslotAPIService', ['create', 'update', 'delete']);

    const mockMatDialogData: SchemaEditDialogData<TimeslotAPIService> = {
      id: '',
      name: '',
      idIsEditable: true,
      serviceForValidation: mockTimeslotAPIService,
      departmentID: 'd_test',
      workplaceID: 'wp_test',
    };
    await TestBed.configureTestingModule({
      imports: [SchemaEditDialogComponent],
      providers: [{ provide: MAT_DIALOG_DATA, useValue: mockMatDialogData }, provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(SchemaEditDialogComponent<TimeslotAPIService>);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should not disable id input if idIsEditable is false', () => {
    expect(component.formGroup.controls.id.disabled).toBeFalse();
  });
});

describe('SchemaEditDialogComponent with WorkplaceServiceValidator', () => {
  let component: SchemaEditDialogComponent<WorkplaceAPIService>;
  let fixture: ComponentFixture<SchemaEditDialogComponent<WorkplaceAPIService>>;
  let mockWorkplaceAPIService: jasmine.SpyObj<WorkplaceAPIService>;

  beforeEach(async () => {
    mockWorkplaceAPIService = jasmine.createSpyObj('WorkplaceAPIService', ['create', 'update', 'delete']);

    const mockMatDialogData: SchemaEditDialogData<WorkplaceAPIService> = {
      id: '',
      name: '',
      idIsEditable: true,
      serviceForValidation: mockWorkplaceAPIService,
      departmentID: 'd_test',
    };
    await TestBed.configureTestingModule({
      imports: [SchemaEditDialogComponent],
      providers: [{ provide: MAT_DIALOG_DATA, useValue: mockMatDialogData }, provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(SchemaEditDialogComponent<WorkplaceAPIService>);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
