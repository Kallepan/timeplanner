import { NgModule } from '@angular/core';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { PlannerStateHandlerService } from './services/planner-state-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';

@NgModule({
  providers: [PersonAPIService, PlannerStateHandlerService, TimetableDataContainerService, PersonDataContainerService, WorkdayAPIService],
})
export class PlannerModule {}
