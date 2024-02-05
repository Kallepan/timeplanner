import { Component, Inject, inject } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { DialogLayoutComponent } from '../dialog-layout/dialog-layout.component';

export type EditTextareaDialogData = {
  title: string;
  label: string;
  placeholder: string;
  hint: string;
  control: FormControl;
};

@Component({
  selector: 'app-edit-textarea-dialog',
  standalone: true,
  imports: [ReactiveFormsModule, MatInputModule, MatFormFieldModule, DialogLayoutComponent],
  templateUrl: './edit-textarea-dialog.component.html',
  styleUrl: './edit-textarea-dialog.component.scss',
})
export class EditTextareaDialogComponent {
  label: string;
  text: string;
  placeholder: string;
  hint: string;
  title: string;

  control: FormControl;

  dialogRef = inject(MatDialogRef<EditTextareaDialogComponent>);

  constructor(@Inject(MAT_DIALOG_DATA) data: EditTextareaDialogData) {
    this.label = data.label;
    this.placeholder = data.placeholder;
    this.hint = data.hint;
    this.title = data.title;

    this.control = data.control;
  }

  handleEditAbortRequest(): void {
    this.dialogRef.close(null);
  }

  handleEditSaveRequest(): void {
    this.dialogRef.close(this.control.value);
  }
}
