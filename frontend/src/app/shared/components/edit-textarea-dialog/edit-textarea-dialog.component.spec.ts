import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormControl } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { EditTextareaDialogComponent, EditTextareaDialogData } from './edit-textarea-dialog.component';

describe('EditTextareaDialogComponent', () => {
  let component: EditTextareaDialogComponent;
  let fixture: ComponentFixture<EditTextareaDialogComponent>;
  let mockDialogData: EditTextareaDialogData;
  let mockDialogRef: jasmine.SpyObj<MatDialogRef<EditTextareaDialogComponent>>;

  beforeEach(async () => {
    mockDialogRef = jasmine.createSpyObj<MatDialogRef<EditTextareaDialogComponent>>('MatDialogRef', ['close']);
    mockDialogData = {
      title: 'Mock Title',
      label: 'Mock Title',
      placeholder: 'Mock Placeholder',
      hint: 'Mock Hint',
      control: new FormControl(''),
    };

    await TestBed.configureTestingModule({
      imports: [EditTextareaDialogComponent],
      providers: [provideNoopAnimations(), { provide: MatDialogRef, useValue: mockDialogRef }, { provide: MAT_DIALOG_DATA, useValue: mockDialogData }],
    }).compileComponents();

    fixture = TestBed.createComponent(EditTextareaDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have mat-form-field', () => {
    const matFormField = fixture.nativeElement.querySelector('mat-form-field');
    expect(matFormField).toBeTruthy();
  });

  it('should have mat-label', () => {
    const matLabel = fixture.nativeElement.querySelector('mat-label');
    expect(matLabel.textContent).toBe('Mock Title');
  });

  it('should have mat-hint', () => {
    const matHint = fixture.nativeElement.querySelector('mat-hint');
    expect(matHint.textContent).toBe('Mock Hint');
  });

  it('should have textarea', () => {
    const textarea = fixture.nativeElement.querySelector('textarea');
    expect(textarea).toBeTruthy();
  });
});
