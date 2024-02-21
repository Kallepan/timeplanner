import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { MatBottomSheet, MatBottomSheetConfig, MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { AbsencyPanelComponent } from '../absency-panel/absency-panel.component';
import { EditableTimetableComponent } from '../editable-timetable/editable-timetable.component';

@Component({
  selector: 'app-planner-landing-page',
  standalone: true,
  imports: [CommonModule, EditableTimetableComponent],
  templateUrl: './planner-landing-page.component.html',
  styleUrl: './planner-landing-page.component.scss',
})
export class PlannerLandingPageComponent {
  timetableDataContainerService = inject(TimetableDataContainerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  plannerStateHandlerService = inject(PlannerStateHandlerService);

  private _bottomSheet = inject(MatBottomSheet);
  _bottomSheetRef: MatBottomSheetRef | null = null;
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

  getLoadingStatus() {
    return this.timetableDataContainerService.isLoading$;
  }
}
