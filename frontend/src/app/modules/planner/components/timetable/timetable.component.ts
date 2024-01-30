import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Timeslot, Weekday, Workplace } from '../../../timetable/interfaces/timetable.interface';
import { SelectPersonComponent } from '../select-person/select-person.component';
import { PersonListComponent } from '../person-list/person-list.component';
import { CdkDropList, CdkDropListGroup } from '@angular/cdk/drag-drop';
import { Person } from '../../../timetable/interfaces/person.interface';
import { ActionsComponent } from '../actions/actions.component';
export type PersonDropIn = {
  person: Person;
  timeslots: Timeslot[];
};

@Component({
  selector: 'app-timetable',
  standalone: true,
  imports: [CommonModule, SelectPersonComponent, PersonListComponent, ActionsComponent, CdkDropList, CdkDropListGroup],
  templateUrl: './timetable.component.html',
  styleUrl: './timetable.component.scss',
})
export class TimetableComponent {
  editable: boolean = !false;
  displayTime: boolean = !false;

  @Input() showActions: boolean = true;
  @Input() weekdayDatas: Weekday[] = [];
  @Input() fullHeight: number = 0;
  @Input() workplaceDatas: Workplace[] = [];

  @Output() personDroppedIn = new EventEmitter<PersonDropIn>();

  protected assignPersonToTimeslot(person: Person, timeslots: Timeslot[]): void {
    // This works because the timeslots are passed by reference. Hooraay JS!
    timeslots.forEach((timeslot) => {
      timeslot.occupiedBy = person;
    });
  }
}
