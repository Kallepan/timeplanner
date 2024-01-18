import { NgModule } from '@angular/core';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';
import { ViewerStateHandlerService } from './services/viewer-state-handler.service';

@NgModule({
  providers: [ViewerStateHandlerService, WorkdayAPIService],
})
export class ViewerModule {}
