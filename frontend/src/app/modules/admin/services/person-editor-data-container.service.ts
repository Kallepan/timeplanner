import { DestroyRef, Injectable, effect, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { constants } from '@app/core/constants/constants';
import { ValidWeekday } from '@app/core/types/weekday';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { catchError, filter, map, of, switchMap, tap, throwError } from 'rxjs';

@Injectable({
  providedIn: null,
})
export class PersonEditorDataContainerService {
  private _activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private _workplaceAPIService = inject(WorkplaceAPIService);
  private _destroyRef$ = inject(DestroyRef);

  private _activePerson = signal<PersonWithMetadata | null>(null);
  set activePerson(person: PersonWithMetadata) {
    this._activePerson.set(person);
  }
  get activePerson$(): PersonWithMetadata | null {
    return this._activePerson();
  }

  private _weekdays = signal<ValidWeekday[]>(constants.POSSIBLE_WEEKDAYS);
  get weekdays$(): ValidWeekday[] {
    return this._weekdays();
  }

  // needs to be fetched from the server
  private _workplaces = signal<WorkplaceWithMetadata[]>([]);
  get workplaces$(): WorkplaceWithMetadata[] {
    return this._workplaces();
  }

  constructor() {
    effect(
      () => {
        of(this._activeDepartmentHandlerService.activeDepartment$)
          .pipe(
            takeUntilDestroyed(this._destroyRef$),
            filter((department): department is string => !!department),
            switchMap((department) =>
              this._workplaceAPIService.getWorkplaces(department).pipe(
                catchError((err) => throwError(() => err)),
                map((resp) => resp.data),
              ),
            ),
            tap((workplaces) => this._workplaces.set(workplaces)),
          )
          .subscribe(() => {
            // do nothing
          });
      },
      { allowSignalWrites: true },
    );
  }
}
