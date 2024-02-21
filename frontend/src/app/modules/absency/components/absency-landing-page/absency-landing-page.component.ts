import { Component, inject } from '@angular/core';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { PersonAutocompleteComponent } from '../person-autocomplete/person-autocomplete.component';
import { YearCalendarComponent } from '../year-calendar/year-calendar.component';
import { PersonPreviewComponent } from '@app/shared/components/person-preview/person-preview.component';

@Component({
  selector: 'app-absency-landing-page',
  standalone: true,
  imports: [PersonAutocompleteComponent, YearCalendarComponent, PersonPreviewComponent],
  templateUrl: './absency-landing-page.component.html',
  styleUrl: './absency-landing-page.component.scss',
})
export class AbsencyLandingPageComponent {
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  activePersonHandlerService = inject(ActivePersonHandlerServiceService);
}
