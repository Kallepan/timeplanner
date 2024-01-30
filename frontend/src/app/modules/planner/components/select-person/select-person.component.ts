import { Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges, inject } from '@angular/core';

import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { map, startWith } from 'rxjs';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { AsyncPipe } from '@angular/common';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-select-person',
  standalone: true,
  imports: [ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatAutocompleteModule, AsyncPipe, MatButtonModule],
  templateUrl: './select-person.component.html',
  styleUrl: './select-person.component.scss',
})
export class SelectPersonComponent implements OnInit, OnChanges {
  personDataContainerService = inject(PersonDataContainerService);

  control = new FormControl<string | PersonWithMetadata>('');
  filteredPersons$ = this.control.valueChanges.pipe(
    startWith(''),
    map((value) => {
      const name = typeof value === 'string' ? value : value?.last_name;
      return name ? this._filter(name) : this.personDataContainerService.persons$;
    }),
  );

  @Output() selected = new EventEmitter<PersonWithMetadata>();
  @Input() selectedPerson: PersonWithMetadata | null = null;

  ngOnInit(): void {
    this.control.setValue(this.selectedPerson);
  }

  displayFn(person: PersonWithMetadata): string {
    return person ? `${person.last_name} (${person.id})` : 'NA';
  }

  private _filter(name: string): PersonWithMetadata[] {
    const filterValue = name.toLowerCase();

    return this.personDataContainerService.persons$.filter((person) => {
      const toBeFiltered = `${person.first_name}${person.last_name}${person.id}`;

      return toBeFiltered.toLowerCase().includes(filterValue);
    });
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['selectedPerson']) {
      this.control.setValue(this.selectedPerson);
    }
  }
}
