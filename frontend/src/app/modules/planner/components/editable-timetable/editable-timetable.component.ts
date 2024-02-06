import { CdkDropList, CdkDropListGroup } from '@angular/cdk/drag-drop';
import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { EditPersonPreviewComponent } from '../edit-person-preview/edit-person-preview.component';
import { PersonListComponent } from '../person-list/person-list.component';
import { SelectPersonComponent } from '../select-person/select-person.component';

@Component({
  selector: 'app-editable-timetable',
  standalone: true,
  imports: [CommonModule, SelectPersonComponent, CdkDropList, CdkDropListGroup, PersonListComponent, EditPersonPreviewComponent],
  templateUrl: './editable-timetable.component.html',
  styleUrl: './editable-timetable.component.scss',
})
export class EditableTimetableComponent {
  plannerStateHandlerService = inject(PlannerStateHandlerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  timetableDataContainerService = inject(TimetableDataContainerService);
  themeHandlerService = inject(ThemeHandlerService);

  personDroppedIntoTimeslotHandler(person: PersonWithMetadata, timeslots: DisplayedWorkdayTimeslot[], actionToBeExecutedOnFailedValidation?: () => void): void {
    timeslots.forEach((timeslot) => {
      this.plannerStateHandlerService.assignPersonToTimeslot(person, timeslot, actionToBeExecutedOnFailedValidation);
    });
  }

  personAssignedToTimeslotEventHandler(person: PersonWithMetadata, timeslot: DisplayedWorkdayTimeslot, actionToBeExecutedOnFailedValidation: () => void): void {
    this.plannerStateHandlerService.assignPersonToTimeslot(person, timeslot, actionToBeExecutedOnFailedValidation);
  }

  personUnassignedFromTimeslotEventHandler(person: PersonWithMetadata, timeslot: DisplayedWorkdayTimeslot): void {
    this.plannerStateHandlerService.unAssignPersonFromTimeslot(person, timeslot);
  }

  weekdays = ['MON', 'TUE', 'WED', 'THU', 'FRI'];
  getSlotsFromMonToFri(slots: DisplayedWorkdayTimeslot[]): DisplayedWorkdayTimeslot[] {
    return slots.filter((s) => this.weekdays.includes(s.weekday));
  }
}
