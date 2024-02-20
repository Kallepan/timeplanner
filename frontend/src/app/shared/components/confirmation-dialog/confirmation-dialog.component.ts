import { Component, Inject } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DialogLayoutComponent } from '../dialog-layout/dialog-layout.component';

export type ConfirmationDialogComponentData = {
  title: string;
  confirmationMessage: string;
};

@Component({
  selector: 'app-confirmation-dialog',
  standalone: true,
  imports: [DialogLayoutComponent],
  template: `
    <app-dialog-layout [control]="dummyControl" [title]="title">
      <div body>
        <p>{{ confirmationMessage }}</p>
      </div>
    </app-dialog-layout>
  `,
  styleUrl: './confirmation-dialog.component.scss',
})
export class ConfirmationDialogComponent {
  title: string;
  confirmationMessage: string;

  // This is a dummy control to make the form valid, it only exists to return a true value upon dialog close
  dummyControl: FormControl = new FormControl(true);

  constructor(@Inject(MAT_DIALOG_DATA) data: ConfirmationDialogComponentData) {
    this.title = data.title;
    this.confirmationMessage = data.confirmationMessage;
  }
}
