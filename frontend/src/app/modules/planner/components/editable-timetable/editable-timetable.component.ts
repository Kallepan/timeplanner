import { CdkDropList, CdkDropListGroup } from '@angular/cdk/drag-drop';
import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Output, inject } from '@angular/core';
import { DisplayedWorkdayTimeslot } from '@app/modules/viewer/interfaces/workplace';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { SelectPersonComponent } from '../select-person/select-person.component';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { PersonListComponent } from '../person-list/person-list.component';
import { EditPersonPreviewComponent } from '../edit-person-preview/edit-person-preview.component';
import { PersonWithMetadata } from '@app/shared/interfaces/person';

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
  @Output() personDroppedIntoTimeslot = new EventEmitter<{
    person: PersonWithMetadata;
    timeslots: DisplayedWorkdayTimeslot[];
  }>();
}
