import { Component, Inject, inject } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { DialogLayoutComponent } from '@app/shared/components/dialog-layout/dialog-layout.component';
import { CreatePerson } from '@app/shared/interfaces/person';
import { MatCheckboxModule } from '@angular/material/checkbox';

export type PersonEditDialogComponentData = Partial<CreatePerson>;

@Component({
  selector: 'app-person-edit-dialog',
  standalone: true,
  imports: [DialogLayoutComponent, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatCheckboxModule],
  templateUrl: './person-edit-dialog.component.html',
  styleUrl: './person-edit-dialog.component.scss',
})
export class PersonEditDialogComponent {
  private _fb = inject(FormBuilder);
  group: FormGroup;

  constructor(@Inject(MAT_DIALOG_DATA) data: PersonEditDialogComponentData) {
    this.group = this._fb.group({
      firstName: [data.first_name ?? '', [Validators.required]],
      lastName: [data.last_name ?? '', [Validators.required]],
      email: [data.email ?? '', [Validators.required, Validators.email]],
      active: [data.active ?? true],
      workingHours: [data.working_hours ?? 8, [Validators.required, Validators.min(1), Validators.max(40), Validators.pattern(/^[0-9]+(\.[0-9]{1,2})?$/)]],
    });

    // disable first name and last name fields if the data is provided
    if (data.first_name || data.last_name) {
      this.group.get('firstName')?.disable();
      this.group.get('lastName')?.disable();
    }
  }
}
