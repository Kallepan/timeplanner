import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges, inject } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatTooltipModule } from '@angular/material/tooltip';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { debounceTime, map, startWith } from 'rxjs';

@Component({
  selector: 'app-select-person',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatAutocompleteModule, MatButtonModule, MatIconModule, MatTooltipModule],
  templateUrl: './select-person.component.html',
  styleUrl: './select-person.component.scss',
})
export class SelectPersonComponent implements OnInit, OnChanges {
  personDataContainerService = inject(PersonDataContainerService);

  @Input() comment: string | null = null;

  control = new FormControl<string | PersonWithMetadata>('');
  filteredPersons$ = this.control.valueChanges.pipe(
    startWith(''),
    debounceTime(150),
    map((value) => {
      const name = typeof value === 'string' ? value : value?.last_name;
      return name ? this._filter(name) : this.personDataContainerService.persons$;
    }),
    map((persons) => persons.filter((person) => person.weekdays?.map((wd) => wd.id).includes(this.weekday))),
  );

  @Output() commentEditRequest = new EventEmitter<void>();
  @Output() selected = new EventEmitter<{ p: PersonWithMetadata; actionToBeExecutedOnFailedValidation?: () => void }>();
  @Input() selectedPerson: PersonWithMetadata | null = null;
  @Input() weekday: string;

  ngOnInit(): void {
    this.control.setValue(this.selectedPerson);
  }

  displayFn(person: PersonWithMetadata): string {
    return person ? `${person.last_name} (${person.id})` : '';
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

  emitEvent(p: PersonWithMetadata): void {
    this.selected.emit({ p, actionToBeExecutedOnFailedValidation: () => this.control.setValue(null) });
  }
}
