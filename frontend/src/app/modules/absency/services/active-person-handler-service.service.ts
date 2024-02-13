import { Injectable, signal } from '@angular/core';
import { PersonWithMetadata } from '@app/shared/interfaces/person';

@Injectable({
  providedIn: null,
})
export class ActivePersonHandlerServiceService {
  private _activePerson = signal<PersonWithMetadata | null>(null);
  set activePerson(person: PersonWithMetadata) {
    this._activePerson.set(person);
  }
  get activePerson$() {
    return this._activePerson();
  }
}
