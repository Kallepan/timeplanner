import { NgModule } from '@angular/core';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { TimetableDataContainerService } from '../../shared/services/timetable-data-container.service';

@NgModule({
  providers: [WorkdayAPIService, TimetableDataContainerService],
})
export class ViewerModule {}
