import { Component, EventEmitter, Input, Output } from '@angular/core';
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
  @Output() toggleTimes = new EventEmitter<boolean>();

  @Input() displayTimes = true;
  @Input() displayColors = true;
}
