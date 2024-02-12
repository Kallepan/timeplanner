import { NgModule } from '@angular/core';
import { PlannerStateHandlerService } from './services/planner-state-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';

@NgModule({
  providers: [PlannerStateHandlerService, TimetableDataContainerService, PersonDataContainerService],
})
export class PlannerModule {}
