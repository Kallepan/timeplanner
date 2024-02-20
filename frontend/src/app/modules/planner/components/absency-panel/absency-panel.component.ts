import { Component, inject } from '@angular/core';
import { MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { AbsencyDataContainerService } from '../../services/absency-data-container.service';
import { CommonModule } from '@angular/common';
import { MatProgressBarModule } from '@angular/material/progress-bar';

@Component({
  selector: 'app-absency-panel',
  standalone: true,
  imports: [CommonModule, MatProgressBarModule],
  providers: [AbsencyDataContainerService],
  templateUrl: './absency-panel.component.html',
  styleUrl: './absency-panel.component.scss',
})
export class AbsencyPanelComponent {
  private _bottomSheetRef = inject(MatBottomSheetRef<AbsencyPanelComponent>);
  absencyDataContainer = inject(AbsencyDataContainerService);

  close(): void {
    this._bottomSheetRef.dismiss();
  }
}
