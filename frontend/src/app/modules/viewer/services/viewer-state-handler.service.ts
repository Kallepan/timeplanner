import { Injectable, inject, signal } from '@angular/core';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';

type ActiveWeek = {
  department: string;
  startDate: Date;
  endDate: Date;
  weekNumber: number;
  year: number;
};

@Injectable({
  providedIn: null,
})
export class ViewerStateHandlerService {
  // inject the services here
  private workdayAPIService = inject(WorkdayAPIService);

  // this keeps track of the active week currently being viewed
  private activeWeek = signal<ActiveWeek | null>(null);
}
