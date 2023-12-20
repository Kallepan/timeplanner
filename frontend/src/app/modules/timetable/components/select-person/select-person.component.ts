import {
  Component,
  EventEmitter,
  Input,
  OnChanges,
  OnInit,
  Output,
  SimpleChanges,
  inject,
} from '@angular/core';
import { PersonDataService } from '../../services/person-data.service';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { Person } from '../../interfaces/person.interface';
import { Observable, map, startWith } from 'rxjs';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'app-select-person',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatAutocompleteModule,
    AsyncPipe,
  ],
  templateUrl: './select-person.component.html',
  styleUrl: './select-person.component.scss',
})
export class SelectPersonComponent implements OnInit, OnChanges {
  private readonly personDataService = inject(PersonDataService);

  control = new FormControl<string | Person>('');
  persons = this.personDataService.persons$();
  filteredPersons: Observable<Person[]> | undefined;

  @Output() selected = new EventEmitter<Person>();
  @Input() selectedPerson: Person | null = null;

  ngOnInit(): void {
    this.filteredPersons = this.control.valueChanges.pipe(
      startWith(''),
      map((value) => {
        const name = typeof value === 'string' ? value : value?.lastname;
        return name ? this._filter(name) : this.persons;
      }),
    );

    this.control.setValue(this.selectedPerson);
  }

  displayFn(person: Person): string {
    return person ? `${person.lastname} ${person.id}` : 'NA';
  }

  private _filter(name: string): Person[] {
    const filterValue = name.toLowerCase();

    return this.persons.filter((person) => {
      const toBeFiltered = `${person.firstname}${person.lastname}${person.id}`;

      return toBeFiltered.toLowerCase().includes(filterValue);
    });
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['selectedPerson']) {
      this.control.setValue(this.selectedPerson);
    }
  }
}
