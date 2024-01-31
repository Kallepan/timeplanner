import { Component, EventEmitter, Output } from '@angular/core';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

@Component({
  selector: 'app-actions-editable',
  standalone: true,
  imports: [MatSlideToggleModule],
  templateUrl: './actions.component.html',
  styleUrl: './actions.component.scss',
})
export class ActionsComponent {
  @Output() toggleColors = new EventEmitter<boolean>();
  @Output() toggleTimeLabel = new EventEmitter<boolean>();
  @Output() togglePersonsLabel = new EventEmitter<boolean>();
}
