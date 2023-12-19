import { Injectable } from '@angular/core';
import { AbstractTimetableDataService } from '@app/modules/timetable/services/timetable-data.service';

@Injectable({
  providedIn: 'root',
})
export class PlannerTimetableDataService extends AbstractTimetableDataService {
  constructor() {
    super();
  }
}
