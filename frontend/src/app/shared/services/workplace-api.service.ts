/**
 * This is the service that handles the API calls for the workplace.
 * It can:
 * - get a workplace
 * - get all workplaces
 * - create a workplace
 * - delete a workplace
 */
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { APIResponse } from '@app/core/interfaces/response';
import { Observable, catchError, map, of } from 'rxjs';
import { CreateWorkplace, WorkplaceWithMetadata } from '../interfaces/workplace';
import { constants } from '@app/core/constants/constants';
import { CheckIDExistsInterface } from '@app/modules/admin/validators/id-validator';

@Injectable({
  providedIn: 'root',
})
export class WorkplaceAPIService implements CheckIDExistsInterface {
  checkIDExists(id: string, departmentID?: string): Observable<boolean> {
    if (departmentID === undefined) {
      return of(true);
    }

    const url = `${constants.APIS.PLANNER}/department/${departmentID}/workplace/${id}`;

    const headerOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<WorkplaceWithMetadata>>(url, headerOptions).pipe(
      map((res) => res.data !== null),
      catchError(() => of(false)),
    );
  }

  private http = inject(HttpClient);

  getWorkplaces(department: string): Observable<APIResponse<WorkplaceWithMetadata[]>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<WorkplaceWithMetadata[]>>(url, httpOptions);
  }

  getWorkplace(department: string, workplace: string): Observable<APIResponse<WorkplaceWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<WorkplaceWithMetadata>>(url, httpOptions);
  }

  createWorkplace(department: string, workplace: CreateWorkplace): Observable<APIResponse<WorkplaceWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.post<APIResponse<WorkplaceWithMetadata>>(url, workplace, httpOptions);
  }

  deleteWorkplace(department: string, workplace: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }
}
