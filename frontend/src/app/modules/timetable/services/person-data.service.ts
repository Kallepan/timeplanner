import { Injectable, WritableSignal, signal } from '@angular/core';
import { Observable, map, of } from 'rxjs';
import { PERSON_DATA } from '../tests/person.data';
import { Person } from '../interfaces/person.interface';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root',
})
export class PersonDataService {
  persons$: WritableSignal<Person[]> = signal([]);

  constructor() {
    this.getPersons()
      .pipe(takeUntilDestroyed())
      .subscribe((persons) => {
        this.persons$.set(persons);
      });
  }

  getPersons(): Observable<Person[]> {
    return of(PERSON_DATA).pipe(
      map((persons) => {
        const formattedPersons = persons.map((person) => ({
          workingHours: person.working_hours,
          actualHours: person.actual_hours,
          fullname: `${person.firstname} ${person.lastname} (${person.id})`,
          ...person,
        }));

        return formattedPersons;
      }),
    );
  }
}
