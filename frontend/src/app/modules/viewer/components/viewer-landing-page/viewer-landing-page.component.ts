import { Component, inject } from '@angular/core';
import { ActionsComponent } from '../actions/actions.component';
import { CommonModule } from '@angular/common';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ViewOnlyTimetableComponent } from '../view-only-timetable/view-only-timetable.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';

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

  getLoadingStatus() {
    return this.timetableDataContainerService.isLoading$;
  }
}
