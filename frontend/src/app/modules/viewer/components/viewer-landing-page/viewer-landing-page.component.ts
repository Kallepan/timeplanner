import { Component, inject } from '@angular/core';
import { ActionsComponent } from '../actions/actions.component';
import { CommonModule } from '@angular/common';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ViewOnlyTimetableComponent } from '../view-only-timetable/view-only-timetable.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { MatBottomSheet, MatBottomSheetConfig, MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { AbsencyPanelComponent } from '@app/modules/planner/components/absency-panel/absency-panel.component';

@Component({
  selector: 'app-viewer-landing-page',
  standalone: true,
  imports: [CommonModule, ViewOnlyTimetableComponent, ActionsComponent],
  templateUrl: './viewer-landing-page.component.html',
  styleUrl: './viewer-landing-page.component.scss',
})
export class ViewerLandingPageComponent {
  // inject the services here
  timetableDataContainerService = inject(TimetableDataContainerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);

  private _bottomSheet = inject(MatBottomSheet);
  _bottomSheetRef: MatBottomSheetRef | null = null;

  getLoadingStatus() {
    return this.timetableDataContainerService.isLoading$;
  }

  toggleAbsencyPanel(): void {
    if (this._bottomSheetRef) {
      this._bottomSheetRef.dismiss();
      this._bottomSheetRef = null;
      return;
    }

    const config: MatBottomSheetConfig = {
      hasBackdrop: false,
      closeOnNavigation: true,
    };
    this._bottomSheetRef = this._bottomSheet.open(AbsencyPanelComponent, config);
  }
}
