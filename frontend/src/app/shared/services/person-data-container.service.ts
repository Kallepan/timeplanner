import { Injectable, effect, inject, signal } from '@angular/core';
import { map } from 'rxjs';
import { PersonWithMetadata } from '../interfaces/person';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';
import { PersonAPIService } from './person-api.service';

/** Service to access the valid persons for a given department */
@Injectable({
  providedIn: null,
})
export class PersonDataContainerService {
  private activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private personAPISerivce = inject(PersonAPIService);

  protected _persons = signal<PersonWithMetadata[]>([]);
  get persons$(): PersonWithMetadata[] {
    return this._persons();
  }

  constructor() {
    effect(
      () => {
        this.personAPISerivce
          .getPersons(this.activeDepartmentHandlerService.activeDepartment$)
          .pipe(map((resp) => resp.data))
          .subscribe({
            next: (persons) => {
              this._persons.set(persons);
            },
          });
      },
      { allowSignalWrites: true },
    );
  }
}
