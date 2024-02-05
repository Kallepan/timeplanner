import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { ActionsComponent } from '../actions/actions.component';
import { EditableTimetableComponent } from '../editable-timetable/editable-timetable.component';
import { PersonListComponent } from '../person-list/person-list.component';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [CommonModule, EditableTimetableComponent, ActionsComponent, PersonListComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.scss',
  providers: [],
})
export class LandingComponent {
  timetableDataContainerService = inject(TimetableDataContainerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  plannerStateHandlerService = inject(PlannerStateHandlerService);
}
