import { Component, Inject, inject } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';

export type EditTextareaDialogData = {
  label: string;
  placeholder: string;
  hint: string;
  control: FormControl;
};

@Component({
  selector: 'app-edit-textarea-dialog',
  standalone: true,
  imports: [ReactiveFormsModule, MatButtonModule, MatInputModule, MatFormFieldModule, MatDialogModule, MatIconModule],
  templateUrl: './edit-textarea-dialog.component.html',
  styleUrl: './edit-textarea-dialog.component.scss',
})
export class EditTextareaDialogComponent {
  label: string;
  text: string;
  placeholder: string;
  hint: string;

  control: FormControl;

  dialogRef = inject(MatDialogRef<EditTextareaDialogComponent>);

  constructor(@Inject(MAT_DIALOG_DATA) data: EditTextareaDialogData) {
    this.label = data.label;
    this.placeholder = data.placeholder;
    this.hint = data.hint;

    this.control = data.control;
  }

  handleEditAbortRequest(): void {
    this.dialogRef.close(null);
  }

  handleEditSaveRequest(): void {
    this.dialogRef.close(this.control.value);
  }
}
