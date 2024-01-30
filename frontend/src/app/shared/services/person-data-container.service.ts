import { Injectable, signal } from '@angular/core';
import { PersonWithMetadata } from '../interfaces/person';

/** Service to access the valid persons for a given department */
@Injectable({
  providedIn: null,
})
export class PersonDataContainerService {
  protected _persons = signal<PersonWithMetadata[]>([]);
  get persons$(): PersonWithMetadata[] {
    return this._persons();
  }
  set persons(value: PersonWithMetadata[]) {
    this._persons.set(value);
  }
}
