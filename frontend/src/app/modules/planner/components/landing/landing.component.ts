import { CommonModule } from '@angular/common';
import { Component, OnInit, inject } from '@angular/core';
import { TimetableComponent } from '@app/modules/timetable/components/timetable/timetable.component';
import { AbstractTimetableDataService } from '@app/modules/timetable/services/timetable-data.service';
import { PlannerTimetableDataService } from '../../services/planner-timetable-data.service';
import { DUMMY_TIMESLOTS } from '../../tests/timeslots.data';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [CommonModule, TimetableComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.scss',
  providers: [
    {
      provide: AbstractTimetableDataService,
      useClass: PlannerTimetableDataService,
    },
  ],
})
export class LandingComponent implements OnInit {
  public calenderWeek: string = 'KW 42';
  public calenderYear: string = '2020';

  public timetableDataService = inject(AbstractTimetableDataService);

  ngOnInit(): void {
    setTimeout(() => {
      this.timetableDataService._slots.next(DUMMY_TIMESLOTS);
      this.timetableDataService.weekdays$ = new Date('2023-01-02');
    }, 0);
  }
}
