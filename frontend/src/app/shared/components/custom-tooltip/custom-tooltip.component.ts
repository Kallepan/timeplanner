import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-custom-tooltip',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="container" [ngStyle]="{ top: topCoordinate, left: leftCoordinate }">
      <div class="content">
        {{ text }}
      </div>
    </div>
  `,
  styleUrl: './custom-tooltip.component.scss',
})
export class CustomTooltipComponent {
  @Input({ required: true }) text: string = '';

  @Input({ required: true }) set mouseX(value: number) {
    this.leftCoordinate = `${value + 10}px`;
  }
  @Input({ required: true }) set mouseY(value: number) {
    this.topCoordinate = `${value + 10}px`;
  }

  topCoordinate: string = '0px'; // e.g. '50px'
  leftCoordinate: string = '0px'; // e.g. '100px'
}
