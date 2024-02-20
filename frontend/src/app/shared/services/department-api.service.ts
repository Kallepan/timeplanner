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
import { Observable } from 'rxjs';
import { DepartmentWithMetadata, Department } from '../interfaces/department';
import { constants } from '@app/constants/constants';
import { AbsenceReponse } from '@app/modules/absency/interfaces/absence';

@Injectable({
  providedIn: 'root',
})
export class DepartmentAPIService {
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

  createDepartment(department: Department): Observable<APIResponse<DepartmentWithMetadata>> {
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
