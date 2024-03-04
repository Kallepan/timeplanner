import { CommonModule } from '@angular/common';
import { Component, Inject, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AsyncIDValidator, CheckIDExistsInterface } from '@app/modules/admin/validators/id-validator';

export type SchemaEditDialogData<T extends CheckIDExistsInterface> = {
  id?: string;
  name?: string;
  idIsEditable: boolean;
  serviceForValidation: T;

  departmentID?: string;
  workplaceID?: string;
};

@Component({
  selector: 'app-create-dialog',
  standalone: true,
  imports: [CommonModule, MatDialogModule, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatButtonModule],
  templateUrl: './schema-edit-dialog.component.html',
  styleUrl: './schema-edit-dialog.component.scss',
})
export class SchemaEditDialogComponent<T extends CheckIDExistsInterface> {
  // Create a form group for the dialog
  private _fb = inject(FormBuilder);
  formGroup = this._fb.group({
    id: ['', [Validators.required, Validators.pattern(/^[a-z0-9]{1,4}$/)]],
    name: ['', [Validators.required, Validators.pattern(/^[a-zA-Z0-9 ]{1,20}$/)]],
  });

  constructor(@Inject(MAT_DIALOG_DATA) public data: SchemaEditDialogData<T>) {
    this.formGroup.controls.id.setValue(data.id ?? '', { emitEvent: false });
    this.formGroup.controls.name.setValue(data.name ?? '', { emitEvent: false });

    this.formGroup.controls.id.addAsyncValidators(AsyncIDValidator(this.data.serviceForValidation, data.departmentID, data.workplaceID));

    if (!data.idIsEditable) {
      this.formGroup.controls.id.disable();
    }
  }
}
