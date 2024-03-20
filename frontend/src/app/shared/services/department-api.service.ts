/**
 * This service handles all interactions with the department API.
 * It can:
 *   - get a department
 *   - get all departments
 *   - create a department
 *   - delete a department
 */
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { APIResponse } from '@app/core/interfaces/response';
import { Observable, catchError, map, of } from 'rxjs';
import { DepartmentWithMetadata, CreateDepartment } from '../interfaces/department';
import { constants } from '@app/core/constants/constants';
import { AbsenceReponse } from '@app/modules/absency/interfaces/absence';
import { CheckIDExistsInterface } from '@app/modules/admin/validators/id-validator';

@Injectable({
  providedIn: 'root',
})
export class DepartmentAPIService implements CheckIDExistsInterface {
  checkIDExists(id: string): Observable<boolean> {
    const url = `${constants.APIS.PLANNER}/department/${id}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<DepartmentWithMetadata>>(url, httpOptions).pipe(
      map((res) => res.data !== null),
      catchError(() => of(false)),
    );
  }
  private http = inject(HttpClient);

  getAbsencesForDepartment(departmentName: string, date: string): Observable<APIResponse<AbsenceReponse[]>> {
    const url = `${constants.APIS.PLANNER}/department/${departmentName}/absency`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      params: {
        date,
      },
    };

    return this.http.get<APIResponse<AbsenceReponse[]>>(url, httpOptions);
  }

  getDepartments(): Observable<APIResponse<DepartmentWithMetadata[]>> {
    const url = `${constants.APIS.PLANNER}/department`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<DepartmentWithMetadata[]>>(url, httpOptions);
  }

  getDepartment(departmentName: string): Observable<APIResponse<DepartmentWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${departmentName}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<DepartmentWithMetadata>>(url, httpOptions);
  }

  createDepartment(department: CreateDepartment): Observable<APIResponse<DepartmentWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.post<APIResponse<DepartmentWithMetadata>>(url, department, httpOptions);
  }

  deleteDepartment(departmentName: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/department/${departmentName}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }
}
