import { CdkDropList, CdkDropListGroup } from '@angular/cdk/drag-drop';
import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { SelectPersonComponent } from '../select-person/select-person.component';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { PersonListComponent } from '../person-list/person-list.component';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { EditPersonPreviewComponent } from '../edit-person-preview/edit-person-preview.component';

@Component({
  selector: 'app-editable-timetable',
  standalone: true,
  imports: [CommonModule, SelectPersonComponent, CdkDropList, CdkDropListGroup, PersonListComponent, EditPersonPreviewComponent],
  templateUrl: './editable-timetable.component.html',
  styleUrl: './editable-timetable.component.scss',
})
export class EditableTimetableComponent {
  plannerStateHandlerService = inject(PlannerStateHandlerService);
  timetableDataContainerService = inject(TimetableDataContainerService);

  personDroppedIntoTimeslotHandler(person: PersonWithMetadata, timeslots: DisplayedWorkdayTimeslot[], actionToBeExecutedOnFailedValidation?: () => void): void {
    this.plannerStateHandlerService.assignPersonToTimelots(person, timeslots, actionToBeExecutedOnFailedValidation);
  }

  weekdays = ['MON', 'TUE', 'WED', 'THU', 'FRI'];
  getSlotsFromMonToFri(slots: DisplayedWorkdayTimeslot[]): DisplayedWorkdayTimeslot[] {
    return slots.filter((s) => this.weekdays.includes(s.weekday));
  }
}
