import { CdkDropList, CdkDropListGroup } from '@angular/cdk/drag-drop';
import { Component, EventEmitter, Output, inject } from '@angular/core';
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
import { ActionsComponent } from '../actions/actions.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { DatePipe, NgClass, NgStyle, TitleCasePipe } from '@angular/common';

@Component({
  selector: 'app-editable-timetable',
  standalone: true,
  imports: [
    DatePipe,
    TitleCasePipe,
    NgStyle,
    NgClass,
    SelectPersonsComponent,
    CdkDropList,
    CdkDropListGroup,
    PersonListComponent,
    PersonPreviewComponent,
    MatIconModule,
    MatButtonModule,
    ActionsComponent,
  ],
  templateUrl: './editable-timetable.component.html',
  styleUrl: './editable-timetable.component.scss',
})
export class EditableTimetableComponent {
  // output
  @Output() toggleAbsencyPanel = new EventEmitter<void>();

  // services
  plannerStateHandlerService = inject(PlannerStateHandlerService);
  activeWeekHandlerService = inject(ActiveWeekHandlerService);
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
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

  weekdays = [1, 2, 3, 4, 5];
  getSlotsFromMonToFri(slots: DisplayedWorkdayTimeslot[]): DisplayedWorkdayTimeslot[] {
    return slots.filter((s) => this.weekdays.includes(s.weekday));
  }
}
