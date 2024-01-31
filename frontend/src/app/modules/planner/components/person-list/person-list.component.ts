/**
 * Component which displays a list of persons to be used for editing the planner.
 */
import { Component, inject } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CdkDrag, CdkDropList } from '@angular/cdk/drag-drop';
import { SearchBarComponent } from '@app/shared/components/search-bar/search-bar.component';
import { FormControl } from '@angular/forms';
import { debounceTime, filter, map, startWith } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { MatProgressBarModule } from '@angular/material/progress-bar';

@Component({
  selector: 'app-person-list',
  standalone: true,
  imports: [SearchBarComponent, MatCardModule, MatTooltipModule, CdkDrag, CdkDropList, AsyncPipe, MatProgressBarModule],
  templateUrl: './person-list.component.html',
  styleUrl: './person-list.component.scss',
})
export class PersonListComponent {
  // data which is displayed in the component
  personDataContainerService = inject(PersonDataContainerService);

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
  );
}
