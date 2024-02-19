import { Component, inject } from '@angular/core';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import { PersonAutocompleteComponent } from '../person-autocomplete/person-autocomplete.component';
import { YearCalendarComponent } from '../year-calendar/year-calendar.component';
import { PersonPreviewComponent } from '@app/shared/components/person-preview/person-preview.component';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [PersonAutocompleteComponent, YearCalendarComponent, PersonPreviewComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.scss',
})
export class LandingComponent {
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  activePersonHandlerService = inject(ActivePersonHandlerServiceService);
}
