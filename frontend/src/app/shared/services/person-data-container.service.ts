import { DestroyRef, Injectable, effect, inject, signal } from '@angular/core';
import { map } from 'rxjs';
import { PersonWithMetadata } from '../interfaces/person';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';
import { PersonAPIService } from './person-api.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

/** Service to access the valid persons for a given department */
@Injectable({
  providedIn: null,
})
export class PersonDataContainerService {
  private destroyRef$ = inject(DestroyRef);
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
          .pipe(
            takeUntilDestroyed(this.destroyRef$),
            map((resp) => resp.data),
          )
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
