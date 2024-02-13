import { Component, inject } from '@angular/core';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import { PersonAutocompleteComponent } from '../person-autocomplete/person-autocomplete.component';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [PersonAutocompleteComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.scss',
})
export class LandingComponent {
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  activePersonHandlerService = inject(ActivePersonHandlerServiceService);
}
