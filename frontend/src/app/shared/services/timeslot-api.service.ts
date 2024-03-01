/**
 * This service is used to make API calls to the backend for timeslot related
 * data.
 * It can:
 * - get a timeslot
 * - get all timeslots
 * - create a timeslot
 * - delete a timeslot
 * - assign a timeslot to a weekday
 * - unassign a timeslot from a weekday
 */
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { APIResponse } from '@app/core/interfaces/response';
import { constants } from '@app/constants/constants';
import { Observable } from 'rxjs';
import { CreateTimeslot, TimeslotWithMetadata } from '../interfaces/timeslot';
@Injectable({
  providedIn: 'root',
})
export class TimeslotAPIService {
  private http = inject(HttpClient);

  getTimeslots(department: string, workplace: string): Observable<APIResponse<TimeslotWithMetadata[]>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<TimeslotWithMetadata[]>>(url, httpOptions);
  }

  getTimeslot(department: string, workplace: string, timeslot: string): Observable<APIResponse<TimeslotWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot/${timeslot}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<TimeslotWithMetadata>>(url, httpOptions);
  }

  createTimeslot(department: string, workplace: string, timeslot: CreateTimeslot): Observable<APIResponse<TimeslotWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.post<APIResponse<TimeslotWithMetadata>>(url, timeslot, httpOptions);
  }

  deleteTimeslot(department: string, workplace: string, timeslot: string): Observable<APIResponse<TimeslotWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot/${timeslot}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.delete<APIResponse<TimeslotWithMetadata>>(url, httpOptions);
  }

  // Relations
  updateTimeslotOnWeekday(department: string, workplace: string, timeslot: string, id: number, start_time: string, end_time: string): Observable<APIResponse<TimeslotWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot/${timeslot}/weekday`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body = {
      id,
      start_time,
      end_time,
    };

    return this.http.put<APIResponse<TimeslotWithMetadata>>(url, body, httpOptions);
  }

  assignTimeslotToWeekday(department: string, workplace: string, timeslot: string, id: number, start_time: string, end_time: string): Observable<APIResponse<TimeslotWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot/${timeslot}/weekday`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body = {
      id,
      start_time,
      end_time,
    };

    return this.http.post<APIResponse<TimeslotWithMetadata>>(url, body, httpOptions);
  }

  unassignTimeslotFromWeekday(department: string, workplace: string, timeslot: string, weekdayID: number): Observable<APIResponse<TimeslotWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/department/${department}/workplace/${workplace}/timeslot/${timeslot}/weekday`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      body: {
        id: weekdayID,
      },
    };

    return this.http.delete<APIResponse<TimeslotWithMetadata>>(url, httpOptions);
  }
}
