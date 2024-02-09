import { CdkDropList, CdkDropListGroup } from '@angular/cdk/drag-drop';
import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { PersonPreviewComponent } from '../../../../shared/components/person-preview/person-preview.component';
import { PersonListComponent } from '../person-list/person-list.component';
import { SelectPersonsComponent } from '../select-persons/select-persons.component';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-editable-timetable',
  standalone: true,
  imports: [CommonModule, SelectPersonsComponent, CdkDropList, CdkDropListGroup, PersonListComponent, PersonPreviewComponent, MatIconModule, MatButtonModule],
  templateUrl: './editable-timetable.component.html',
  styleUrl: './editable-timetable.component.scss',
})
export class EditableTimetableComponent {
  plannerStateHandlerService = inject(PlannerStateHandlerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  timetableDataContainerService = inject(TimetableDataContainerService);
  themeHandlerService = inject(ThemeHandlerService);

  personDroppedIntoTimeslotHandler(person: PersonWithMetadata, timeslots: DisplayedWorkdayTimeslot[]): void {
    timeslots.forEach((timeslot) => {
      this.plannerStateHandlerService.assignPersonToTimeslot(person, timeslot);
    });
  }

  personAssignedToTimeslotEventHandler(person: PersonWithMetadata, timeslot: DisplayedWorkdayTimeslot): void {
    this.plannerStateHandlerService.assignPersonToTimeslot(person, timeslot);
  }

  personUnassignedFromTimeslotEventHandler(person: PersonWithMetadata, timeslot: DisplayedWorkdayTimeslot): void {
    this.plannerStateHandlerService.unAssignPersonFromTimeslot(person, timeslot);
  }

  weekdays = ['MON', 'TUE', 'WED', 'THU', 'FRI'];
  getSlotsFromMonToFri(slots: DisplayedWorkdayTimeslot[]): DisplayedWorkdayTimeslot[] {
    return slots.filter((s) => this.weekdays.includes(s.weekday));
  }
}
