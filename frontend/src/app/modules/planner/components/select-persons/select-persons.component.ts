import { Component, ElementRef, EventEmitter, Input, Output, ViewChild, inject } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatChipInputEvent, MatChipsModule } from '@angular/material/chips';
import { MatAutocompleteModule, MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';
import { debounceTime, filter, map } from 'rxjs';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { ENTER, COMMA } from '@angular/cdk/keycodes';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { AsyncPipe } from '@angular/common';
@Component({
  selector: 'app-select-persons',
  standalone: true,
  imports: [AsyncPipe, MatIconModule, ReactiveFormsModule, MatButtonModule, MatFormFieldModule, MatInputModule, MatChipsModule, MatAutocompleteModule],
  templateUrl: './select-persons.component.html',
  styleUrl: './select-persons.component.scss',
})
export class SelectPersonsComponent {
  seperatorKeysCode: number[] = [ENTER, COMMA];

  personDataContainerService = inject(PersonDataContainerService);
  // initially
  @Input() initiallySetPersons: PersonWithMetadata[] = [];
  @Input() comment: string | null = null;

  @ViewChild('personInput') personInput: ElementRef<HTMLInputElement>;
  control = new FormControl<string>('');
  filteredPersons$ = this.control.valueChanges.pipe(
    debounceTime(150),
    filter((value) => typeof value === 'string'),
    map((value) => this._filter(value!)),
    // Filter out the persons that are normally not available on the selected weekday
    // I disabled this for now, because this feature is not wished by the users
    // map((persons) => persons.filter((person) => person.weekdays?.map((wd) => wd.id).includes(this.weekday))),
  );

  @Input() initiallySelectedPersonFromParent: PersonWithMetadata[] = [];
  @Input() weekday: number;

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

  add(event: MatChipInputEvent): void {
    const value = (event.value || '').trim();

    // get the person from the list of persons
    const person = this.personDataContainerService.persons$.find((p) => p.id === value);

    if (!person) return;

    this.emitPersonAssignedToTimeslotEvent(person);

    // Clear the input value
    event.chipInput!.clear();
    this.control.setValue(null);
  }

  remove(person: PersonWithMetadata): void {
    this.emitPersonUnassignedFromTimeslotEvent(person);
  }

  selected(event: MatAutocompleteSelectedEvent): void {
    this.emitPersonAssignedToTimeslotEvent(event.option.value);
    this.personInput.nativeElement.value = '';
    this.control.setValue('', { emitEvent: false });
  }

  @Output() commentDeleteRequest = new EventEmitter<void>();
  @Output() commentEditRequest = new EventEmitter<void>();
  @Output() personAssignedToTimeslot = new EventEmitter<{ p: PersonWithMetadata }>();
  @Output() personUnassignedFromTimeslot = new EventEmitter<{ p: PersonWithMetadata }>();
  emitPersonAssignedToTimeslotEvent(p: PersonWithMetadata): void {
    this.personAssignedToTimeslot.emit({ p });
  }

  emitPersonUnassignedFromTimeslotEvent(p: PersonWithMetadata): void {
    this.personUnassignedFromTimeslot.emit({ p });
  }
}
