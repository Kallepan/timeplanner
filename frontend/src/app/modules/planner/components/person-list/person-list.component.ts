/**
 * Component which displays a list of persons to be used for editing the planner.
 */
import { Component, inject } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CdkDrag, CdkDropList } from '@angular/cdk/drag-drop';
import { SearchBarComponent } from '@app/shared/components/search-bar/search-bar.component';
import { FormControl } from '@angular/forms';
import { combineLatestWith, debounceTime, filter, map, startWith } from 'rxjs';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { toObservable } from '@angular/core/rxjs-interop';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'app-person-list',
  standalone: true,
  imports: [AsyncPipe, SearchBarComponent, MatCardModule, MatTooltipModule, CdkDrag, CdkDropList, MatProgressBarModule],
  templateUrl: './person-list.component.html',
  styleUrl: './person-list.component.scss',
})
export class PersonListComponent {
  // data which is displayed in the component
  personDataContainerService = inject(PersonDataContainerService);
  timetableDataContainerService = inject(TimetableDataContainerService);

  // Filter control
  control = new FormControl<string>('');

  filteredPersons$ = this.control.valueChanges.pipe(
    startWith(''),
    debounceTime(300),
    filter((value): value is string => typeof value === 'string'),
    map((value) => value.toLowerCase()),
    map((value) =>
      this.personDataContainerService.persons$.filter((person) => person.id.includes(value) || person.first_name.toLowerCase().includes(value) || person.last_name.toLowerCase().includes(value)),
    ),
    combineLatestWith(toObservable(this.timetableDataContainerService.listOfPersonsAssignedToTheWholeWeek)),
    map(([persons, listOfPersonsAssignedToTheWholeWeek]) => {
      // filter out the persons whose id is already in the list of persons assigned to the whole week
      return persons.filter((person) => !listOfPersonsAssignedToTheWholeWeek.includes(person.id));
    }),
  );
}
