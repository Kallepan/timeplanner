/**
 * This component is used to display the timetable in a read-only mode.
 * It is used in the view-only-timetable module. To use it import this component,
 * inject the associated service and use it in your template. The service is used
 * to get format and handle the data from the backend.
 **/
import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ActionsComponent } from '../actions/actions.component';

@Component({
  selector: 'app-view-only-timetable',
  standalone: true,
  imports: [CommonModule, ActionsComponent],
  templateUrl: './view-only-timetable.component.html',
  styleUrl: './view-only-timetable.component.scss',
})
export class ViewOnlyTimetableComponent {
  timetableDataContainerService = inject(TimetableDataContainerService);
  activeWeekdayService = inject(ActiveWeekHandlerService);
  themeHandlerService = inject(ThemeHandlerService);
}
