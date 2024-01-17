import { NgModule } from '@angular/core';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';

@NgModule({
  providers: [WorkdayAPIService],
})
export class ViewerModule {}
