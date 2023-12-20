import { Component, EventEmitter, Output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

@Component({
  selector: 'app-actions',
  standalone: true,
  imports: [MatButtonModule, MatSlideToggleModule],
  templateUrl: './actions.component.html',
  styleUrl: './actions.component.scss',
})
export class ActionsComponent {
  @Output() toggleEditing = new EventEmitter<boolean>();
  @Output() toggleTimes = new EventEmitter<boolean>();
}
