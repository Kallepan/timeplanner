import { AfterViewInit, Component, ElementRef, ViewChild, ViewEncapsulation, inject } from '@angular/core';
import Calendar from 'js-year-calendar';
import CalendarDataSourceElement from 'js-year-calendar/dist/interfaces/CalendarDataSourceElement';
import CalendarDayEventObject from 'js-year-calendar/dist/interfaces/CalendarDayEventObject';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';

@Component({
  selector: 'app-year-calendar',
  standalone: true,
  imports: [],
  templateUrl: './year-calendar.component.html',
  styleUrl: './year-calendar.component.scss',
  encapsulation: ViewEncapsulation.ShadowDom,
})
export class YearCalendarComponent implements AfterViewInit {
  private activePersonHandlerService = inject(ActivePersonHandlerServiceService);

  private calendar: Calendar<CalendarDataSourceElement>;
  @ViewChild('calendar', { static: true }) calendarElement: ElementRef | undefined;

  ngAfterViewInit(): void {
    if (!this.calendarElement) {
      return;
    }
    this.calendar = new Calendar(this.calendarElement.nativeElement, {
      language: 'de',
      loadingTemplate: `<div></div>`,
      style: 'background',
      clickDay: (e: CalendarDayEventObject<CalendarDataSourceElement>) => {
        console.log(e);
      },
    });
  }
}
