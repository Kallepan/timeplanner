/**
 *
 * Service to handle active department. It can be set from code or from router query params.
 */
import { Injectable, WritableSignal, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { ActivatedRoute } from '@angular/router';
import { filter, map } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ActiveDepartmentHandlerService {
  // get department from router query params
  private activatedRoute = inject(ActivatedRoute);

  private _activeDepartment: WritableSignal<string | null>;

  // or set department from code
  set activeDepartment(department: string | null) {
    if (!department) {
      this._activeDepartment.set(null);
      return;
    }

    this._activeDepartment.set(this._formatDepartment(department));
  }
  get activeDepartment$() {
    return this._activeDepartment();
  }

  private _formatDepartment(department: string) {
    /* Ensure department is in lowercase and more if needed*/
    return department.toLowerCase();
  }

  constructor() {
    this.activatedRoute.queryParams
      .pipe(
        takeUntilDestroyed(),
        map((params) => params['department']),
        filter((department): department is string => typeof department === 'string'),
        map((department) => this._formatDepartment(department)),
      )
      .subscribe((departmentID) => {
        this._activeDepartment = signal(departmentID);
      });
  }
}
