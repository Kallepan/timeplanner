import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { Weekday, Workplace } from '../../interfaces/timetable.interface';
import { SelectPersonForTimeslotComponent } from '../select-person-for-timeslot/select-person-for-timeslot.component';

@Component({
  selector: 'app-timetable',
  standalone: true,
  imports: [CommonModule, SelectPersonForTimeslotComponent],
  templateUrl: './timetable.component.html',
  styleUrl: './timetable.component.scss',
})
export class TimetableComponent {
  @Input() displayTime: boolean = true;
  @Input() weekdayDatas: Weekday[] = [];
  @Input() fullHeight: number = 0;
  @Input() workplaceDatas: Workplace[] = [];
}
