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
import {
  DetailedDepartmentWithMetadata,
  Department,
} from '../interfaces/department';
import { constants } from '@app/constants/constants';

@Injectable({
  providedIn: 'root',
})
export class DepartmentAPIService {
  private http = inject(HttpClient);

  getDepartments(): Observable<APIResponse<DetailedDepartmentWithMetadata[]>> {
    const url = `${constants.APIS.PLANNER}/department`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<DetailedDepartmentWithMetadata[]>>(
      url,
      httpOptions,
    );
  }

  getDepartment(
    departmentName: string,
  ): Observable<APIResponse<DetailedDepartmentWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${departmentName}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<DetailedDepartmentWithMetadata>>(
      url,
      httpOptions,
    );
  }

  createDepartment(
    department: Department,
  ): Observable<APIResponse<DetailedDepartmentWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.post<APIResponse<DetailedDepartmentWithMetadata>>(
      url,
      department,
      httpOptions,
    );
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
