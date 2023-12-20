import { Component, inject } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { PersonDataService } from '../../services/person-data.service';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CdkDrag, CdkDropList } from '@angular/cdk/drag-drop';
import { SearchBarComponent } from '@app/shared/components/search-bar/search-bar.component';
import { FormControl } from '@angular/forms';
import { filter, map, startWith } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'app-person-list',
  standalone: true,
  imports: [
    SearchBarComponent,
    MatCardModule,
    MatTooltipModule,
    CdkDrag,
    CdkDropList,
    AsyncPipe,
  ],
  templateUrl: './person-list.component.html',
  styleUrl: './person-list.component.scss',
})
export class PersonListComponent {
  // Filter control
  control = new FormControl<string>('');

  // data which is displayed in the component
  private readonly personDataService = inject(PersonDataService);
  persons = this.personDataService.persons$();
  filteredPersons$ = this.control.valueChanges.pipe(
    startWith(''),
    filter((value): value is string => typeof value === 'string'),
    map((value) => value.toLowerCase()),
    map((value) =>
      this.persons.filter((person) =>
        person.fullname.toLowerCase().includes(value),
      ),
    ),
  );
}
