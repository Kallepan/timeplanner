import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonEditDialogComponent, PersonEditDialogComponentData } from './person-edit-dialog.component';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { provideNoopAnimations } from '@angular/platform-browser/animations';

describe('PersonEditDialogComponent with empty person', () => {
  let component: PersonEditDialogComponent;
  let fixture: ComponentFixture<PersonEditDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PersonEditDialogComponent],
      providers: [
        {
          provide: MAT_DIALOG_DATA,
          useValue: {},
        },
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonEditDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

describe('PersonEditDialogComponent with valid person', () => {
  let component: PersonEditDialogComponent;
  let fixture: ComponentFixture<PersonEditDialogComponent>;

  const dialogData: PersonEditDialogComponentData = {
    first_name: 'John',
    last_name: 'Doe',
    email: 'test@example.com',
    working_hours: 8,
    active: true,
  };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PersonEditDialogComponent],
      providers: [
        {
          provide: MAT_DIALOG_DATA,
          useValue: dialogData,
        },
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonEditDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
