import { Component, inject } from '@angular/core';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { PersonAutocompleteComponent } from '../person-autocomplete/person-autocomplete.component';
import { YearCalendarComponent } from '../year-calendar/year-calendar.component';
import { PersonPreviewComponent } from '@app/shared/components/person-preview/person-preview.component';
import { constants } from '@app/core/constants/constants';
import { MatListModule } from '@angular/material/list';
import { NgStyle } from '@angular/common';

@Component({
  selector: 'app-absency-landing-page',
  standalone: true,
  imports: [NgStyle, PersonAutocompleteComponent, YearCalendarComponent, PersonPreviewComponent, MatListModule],
  templateUrl: './absency-landing-page.component.html',
  styleUrl: './absency-landing-page.component.scss',
})
export class AbsencyLandingPageComponent {
  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  activePersonHandlerService = inject(ActivePersonHandlerServiceService);

  absencyReasons = constants.ABSENCY_REASONS;
}
