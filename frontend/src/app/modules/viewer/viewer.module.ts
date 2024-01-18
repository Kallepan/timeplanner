import { NgModule } from '@angular/core';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { ViewerStateHandlerService } from './services/viewer-state-handler.service';
import { TimetableDataContainerService } from './services/timetable-data-container.service';

@NgModule({
  providers: [
    ViewerStateHandlerService,
    WorkdayAPIService,
    TimetableDataContainerService,
  ],
})
export class ViewerModule {}
