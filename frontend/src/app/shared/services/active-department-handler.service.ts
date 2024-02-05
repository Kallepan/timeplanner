/**
 *
 * Service to handle active department. It can be set from code or from router query params.
 * Upon
 */
import { Injectable, inject, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Injectable({
  providedIn: 'root',
})
export class ActiveDepartmentHandlerService {
  // get department from router query params
  private activatedRoute = inject(ActivatedRoute);

  private _activeDepartment = signal<string>(this._formatDepartment(this.activatedRoute.snapshot.queryParams['department']));

  // or set department from code
  set activeDepartment(department: string | undefined) {
    this._activeDepartment.set(this._formatDepartment(department));
  }
  get activeDepartment$() {
    return this._activeDepartment();
  }

  private _formatDepartment(department: string | undefined) {
    /* Ensure department is in lowercase and more if needed*/
    if (!department) {
      return '';
    }
    return department.toLowerCase();
  }
}
