import { Component, Inject } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DialogLayoutComponent } from '@app/shared/components/dialog-layout/dialog-layout.component';

export type DeleteAbsencyDialogComponentData = {
  personID: string;
  date: Date;
};

@Component({
  selector: 'app-delete-absency-dialog',
  standalone: true,
  imports: [DialogLayoutComponent],
  templateUrl: './delete-absency-dialog.component.html',
  styleUrl: './delete-absency-dialog.component.scss',
})
export class DeleteAbsencyDialogComponent {
  title: string;
  date: Date;

  // This is a dummy control to make the form valid, it only exists to return a true value upon dialog close
  dummyControl: FormControl = new FormControl(true);

  constructor(@Inject(MAT_DIALOG_DATA) data: DeleteAbsencyDialogComponentData) {
    this.date = data.date;
    this.title = `Abwesenheit löschen für ${data.personID}`;
  }
}
