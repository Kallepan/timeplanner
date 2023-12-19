import { CommonModule, DatePipe, TitleCasePipe } from '@angular/common';
import { Component, Input } from '@angular/core';
import { Weekday, Workplace } from '../interfaces/timetable.interface';

@Component({
  selector: 'app-timetable',
  standalone: true,
  imports: [CommonModule, DatePipe, TitleCasePipe],
  templateUrl: './timetable.component.html',
  styleUrl: './timetable.component.scss',
})
export class TimetableComponent {
  @Input() displayTime: boolean = true;
  @Input() weekdayDatas: Weekday[] = [];
  @Input() fullHeight: number = 0;
  @Input() workplaceDatas: Workplace[] = [];
}
