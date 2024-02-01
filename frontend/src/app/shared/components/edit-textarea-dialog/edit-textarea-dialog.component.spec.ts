import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { EditTextareaDialogComponent, EditTextareaDialogData } from './edit-textarea-dialog.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { FormControl } from '@angular/forms';

describe('EditTextareaDialogComponent', () => {
  let component: EditTextareaDialogComponent;
  let fixture: ComponentFixture<EditTextareaDialogComponent>;
  let mockDialogData: EditTextareaDialogData;
  let mockDialogRef: jasmine.SpyObj<MatDialogRef<EditTextareaDialogComponent>>;

  beforeEach(async () => {
    mockDialogRef = jasmine.createSpyObj<MatDialogRef<EditTextareaDialogComponent>>('MatDialogRef', ['close']);
    mockDialogData = {
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

  it('save-button should close the dialog with the new text', () => {
    // set control value
    component.control.setValue('New Text');
    fixture.detectChanges();

    // fetch button by id #save-button
    const button = fixture.nativeElement.querySelector('#save-button');

    // click button
    button.click();

    // assert dialog close
    expect(mockDialogRef.close).toHaveBeenCalledWith('New Text');
  });

  it('abort-button should close the dialog with null', () => {
    // fetch button by id #abort-button
    const button = fixture.nativeElement.querySelector('#abort-button');

    // click button
    button.click();

    // assert dialog close
    expect(mockDialogRef.close).toHaveBeenCalledWith(null);
  });

  it('save-button should be disabled if control is invalid', () => {
    // set control value
    component.control.setErrors({ required: true });
    fixture.detectChanges();

    // fetch button by id #save-button
    const button = fixture.nativeElement.querySelector('#save-button');

    // assert button is disabled
    expect(button.disabled).toBeTrue();
  });
});
