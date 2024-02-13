/**
 * Actions component
 *
 * Allows the selection of a person using an mat autocomplete field
 */
import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Output, inject } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { debounceTime, filter, map, startWith } from 'rxjs';

@Component({
  selector: 'app-person-autocomplete',
  standalone: true,
  imports: [MatAutocompleteModule, MatInputModule, MatFormFieldModule, ReactiveFormsModule, CommonModule],
  templateUrl: './person-autocomplete.component.html',
  styleUrl: './person-autocomplete.component.scss',
})
export class PersonAutocompleteComponent {
  personDataContainerService = inject(PersonDataContainerService);

  @Output() personSelected = new EventEmitter<PersonWithMetadata>();
  personControl = new FormControl<string>('');
  filteredPersons$ = this.personControl.valueChanges.pipe(
    startWith(''),
    debounceTime(150),
    filter((value) => typeof value === 'string'),
    map((value) => this._filter(value!)),
  );

  displayFn(person: PersonWithMetadata | null): string {
    return person ? `${person.last_name} (${person.id})` : '';
  }

  private _filter(value: string): PersonWithMetadata[] {
    const filterValue = value.toLowerCase();

    return this.personDataContainerService.persons$.filter((person) => {
      const toBeFiltered = `${person.first_name}${person.last_name}${person.id}`;

      return toBeFiltered.toLowerCase().includes(filterValue);
    });
  }

  // only for testing purposes
  filter(value: string): PersonWithMetadata[] {
    const filterValue = value.toLowerCase();

    return this.personDataContainerService.persons$.filter((person) => {
      const toBeFiltered = `${person.first_name}${person.last_name}${person.id}`;

      return toBeFiltered.toLowerCase().includes(filterValue);
    });
  }
}
