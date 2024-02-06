/**
 * This service handles all interactions with the workday API.
 * It can:
 *  - get all workdays for a department on a given date
 *  - get a specific workday using department, date, workplace, and timeslot
 *  - assign a person to a workday
 *  - remove/unassign a person from a workday
 */
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';
import { constants } from '@app/constants/constants';
import { APIResponse } from '@app/core/interfaces/response';
import { AssignPersonToWorkdayTimeslotRequest, UnassignPersonFromWorkdayTimeslotRequest, UpdateWorkdayRequest, WorkdayTimeslot } from '../interfaces/workday_timeslot';

@Injectable({
  providedIn: null,
})
export class WorkdayAPIService {
  private http = inject(HttpClient);

  getDetailWorkday(departmentName: string, date: string, workplace: string, timeslot: string): Observable<APIResponse<WorkdayTimeslot>> {
    const url = `${constants.APIS.PLANNER}/workday/detail`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      params: new HttpParams({
        fromObject: {
          department: departmentName,
          date: date,
          workplace: workplace,
          timeslot: timeslot,
        },
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<WorkdayTimeslot>>(url, httpOptions);
  }

  getWorkdays(departmentName: string, date: string): Observable<APIResponse<WorkdayTimeslot[]>> {
    const url = `${constants.APIS.PLANNER}/workday`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      params: new HttpParams({
        fromObject: {
          department: departmentName,
          date: date,
        },
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<WorkdayTimeslot[]>>(url, httpOptions);
  }

  updateWorkday(updateWorkdayRequest: UpdateWorkdayRequest): Observable<APIResponse<WorkdayTimeslot>> {
    const url = `${constants.APIS.PLANNER}/workday`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      params: new HttpParams({
        fromObject: {
          departmentID: updateWorkdayRequest.department_id,
          date: updateWorkdayRequest.date,
          workplaceID: updateWorkdayRequest.workplace_id,
          timeslotID: updateWorkdayRequest.timeslot_id,
        },
      }),
    };

    const body = {
      comment: updateWorkdayRequest.comment,
      start_time: updateWorkdayRequest.start_time,
      end_time: updateWorkdayRequest.end_time,
      active: updateWorkdayRequest.active,
    };

    return this.http.put<APIResponse<WorkdayTimeslot>>(url, body, httpOptions);
  }

  assignPerson(department_id: string, date: string, workplace_id: string, timeslot_id: string, person_id: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/workday/assign`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body: AssignPersonToWorkdayTimeslotRequest = {
      department_id,
      date,
      workplace_id,
      timeslot_id,
      person_id,
    };

    return this.http.post<APIResponse<null>>(url, body, httpOptions);
  }

  unassignPerson(department_id: string, date: string, workplace_id: string, timeslot_id: string, person_id: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/workday/assign`;

    const body: UnassignPersonFromWorkdayTimeslotRequest = {
      department_id,
      date,
      workplace_id,
      timeslot_id,
      person_id,
    };

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      body,
    };

    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }
}
