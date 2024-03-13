import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-actions-viewer',
  standalone: true,
  imports: [MatSlideToggleModule, MatButtonModule, RouterLink],
  templateUrl: './actions.component.html',
  styleUrl: './actions.component.scss',
})
export class ActionsComponent {
  @Input() displayComments = false;
  @Input() displayTimes = true;
  @Input() displayColors = true;
  @Input() displayColorsMissing = true;

  @Output() toggleComments = new EventEmitter<boolean>();
  @Output() toggleColors = new EventEmitter<boolean>();
  @Output() toggleTimeLabel = new EventEmitter<boolean>();
  @Output() toggleColorsMissing = new EventEmitter<boolean>();
  @Output() toggleAbsencyPanel = new EventEmitter<boolean>();

  @Output() shiftWeek = new EventEmitter<number>();

  @Input() departmentId: string;
}
