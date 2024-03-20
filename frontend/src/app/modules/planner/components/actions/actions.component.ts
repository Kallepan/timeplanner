import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-actions-editable',
  standalone: true,
  imports: [MatSlideToggleModule, MatButtonModule, RouterLink],
  templateUrl: './actions.component.html',
  styleUrl: './actions.component.scss',
})
export class ActionsComponent {
  @Output() toggleColors = new EventEmitter<boolean>();
  @Output() toggleTimes = new EventEmitter<boolean>();

  @Input() displayTimes = true;
  @Input() displayColors = true;

  @Output() shiftWeek = new EventEmitter<number>();

  @Input() departmentId: string | null;

  @Output() toggleAbsencyPanel = new EventEmitter<void>();
}
