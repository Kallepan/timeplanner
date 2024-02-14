import { Component, OnInit, inject } from '@angular/core';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import { PersonAutocompleteComponent } from '../person-autocomplete/person-autocomplete.component';
import { YearCalendarComponent } from '../year-calendar/year-calendar.component';
import { MatCardModule } from '@angular/material/card';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [PersonAutocompleteComponent, YearCalendarComponent, MatCardModule],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.scss',
})
export class LandingComponent implements OnInit {
  ngOnInit(): void {
    // debug code setup some active person
    this.activePersonHandlerService.activePerson = {
      id: '1',
      last_name: 'Doe',
      first_name: 'John',
      email: 'test@example.com',
      active: true,
      working_hours: 40,
      workplaces: [],
      departments: [],
      weekdays: [],

      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    };
  }

  activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  activePersonHandlerService = inject(ActivePersonHandlerServiceService);
}
