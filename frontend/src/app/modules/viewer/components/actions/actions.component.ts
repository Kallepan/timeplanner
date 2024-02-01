import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

@Component({
  selector: 'app-actions-viewer',
  standalone: true,
  imports: [MatSlideToggleModule],
  templateUrl: './actions.component.html',
  styleUrl: './actions.component.scss',
})
export class ActionsComponent {
  @Input() displayComments = false;
  @Input() displayPersons = true;
  @Input() displayTimes = true;
  @Input() displayColors = true;

  @Output() toggleComments = new EventEmitter<boolean>();
  @Output() toggleColors = new EventEmitter<boolean>();
  @Output() toggleTimeLabel = new EventEmitter<boolean>();
  @Output() togglePersonsLabel = new EventEmitter<boolean>();
}
