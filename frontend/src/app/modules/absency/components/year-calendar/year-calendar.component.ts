import { AfterViewInit, Component, ElementRef, ViewChild, ViewEncapsulation, effect, inject } from '@angular/core';
import Calendar from 'js-year-calendar';
import CalendarDataSourceElement from 'js-year-calendar/dist/interfaces/CalendarDataSourceElement';
import CalendarDayEventObject from 'js-year-calendar/dist/interfaces/CalendarDayEventObject';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import CalendarYearChangedEventObject from 'js-year-calendar/dist/interfaces/CalendarYearChangedEventObject';
import { CustomTooltipComponent } from '@app/shared/components/custom-tooltip/custom-tooltip.component';

@Component({
  selector: 'app-year-calendar',
  standalone: true,
  imports: [CustomTooltipComponent],
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
      loadingTemplate: `<div>Lädt</div>`,
      style: 'background',
      weekStart: 1,
      disabledWeekDays: [],
      clickDay: (e: CalendarDayEventObject<CalendarDataSourceElement>) => {
        this.activePersonHandlerService.handleDayClick(e);
      },
      yearChanged: (year: CalendarYearChangedEventObject) => {
        this.activePersonHandlerService.activeYear = year.currentYear;
      },
      mouseOnDay: (e: CalendarDayEventObject<CalendarDataSourceElement>) => {
        this.handleMouseOnDay(e);
      },
      mouseOutDay: () => {
        this.handleMouseOutDay();
      },
    });
  }

  tooltipText = '';
  tooltipVisible = false;
  mouseTrackerX = 0;
  mouseTrackerY = 0;
  private handleMouseOnDay(e: CalendarDayEventObject<CalendarDataSourceElement>) {
    if (!e.events.length) return;

    this.tooltipText = e.events
      .map((event) => {
        return `${event.name} - ${e.date.toLocaleDateString()}`;
      })
      .join('<br>');
    this.tooltipVisible = true;

    // fetch mouse position from document
  }

  private handleMouseOutDay() {
    this.tooltipVisible = false;
    this.tooltipText = '';
  }

  constructor() {
    effect(() => {
      // Use the absences$ signal to update the calendar
      const fetchedAbsences = this.activePersonHandlerService.absences$;

      this.calendar.setDataSource(fetchedAbsences, true);
      this.calendar.render();
    });
  }
}
