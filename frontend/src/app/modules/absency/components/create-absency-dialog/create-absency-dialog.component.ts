import { Component, Inject, inject } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { DialogLayoutComponent } from '@app/shared/components/dialog-layout/dialog-layout.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatRadioModule } from '@angular/material/radio';
import { constants } from '@app/constants/constants';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';

export type CreateAbsencyDialogData = {
  personID: string;
  startDate: Date;
};

@Component({
  selector: 'app-create-absency-dialog',
  standalone: true,
  imports: [CommonModule, DialogLayoutComponent, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatDatepickerModule, MatRadioModule],
  templateUrl: './create-absency-dialog.component.html',
  styleUrl: './create-absency-dialog.component.scss',
})
export class CreateAbsencyDialogComponent {
  absencyReasons = constants.ABSENCY_REASONS;

  fb = inject(FormBuilder);
  control: FormGroup;
  title: string;
  startDate: Date;

  constructor(@Inject(MAT_DIALOG_DATA) data: CreateAbsencyDialogData) {
    this.startDate = data.startDate;
    this.title = `Abwesenheit erstellen für ${data.personID}`;

    this.control = this.fb.group({
      endDate: [this.startDate, [Validators.required]],
      reason: ['', [Validators.required]],
    });
  }
}
