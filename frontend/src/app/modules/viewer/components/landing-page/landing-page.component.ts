import { AsyncPipe, CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { ViewOnlyTimetableComponent } from '@app/modules/viewer/components/view-only-timetable/view-only-timetable.component';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ActionsComponent } from '../actions/actions.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';

@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [CommonModule, MatButtonModule, RouterLink, AsyncPipe, ViewOnlyTimetableComponent, ActionsComponent],
  templateUrl: './landing-page.component.html',
  styleUrl: './landing-page.component.scss',
})
export class LandingPageComponent {
  // inject the services here
  timetableDataContainerService = inject(TimetableDataContainerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
}
