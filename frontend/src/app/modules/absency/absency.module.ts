import { NgModule } from '@angular/core';
import { ActivePersonHandlerServiceService } from './services/active-person-handler-service.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';

@NgModule({
  providers: [ActivePersonHandlerServiceService, PersonDataContainerService],
})
export class AbsencyModule {}
