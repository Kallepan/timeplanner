import { Component, inject } from '@angular/core';
import { MatBottomSheetRef } from '@angular/material/bottom-sheet';

@Component({
  selector: 'app-absency-panel',
  standalone: true,
  imports: [],
  templateUrl: './absency-panel.component.html',
  styleUrl: './absency-panel.component.scss',
})
export class AbsencyPanelComponent {
  private _bottomSheetRef = inject(MatBottomSheetRef<AbsencyPanelComponent>);

  close(): void {
    this._bottomSheetRef.dismiss();
  }
}
