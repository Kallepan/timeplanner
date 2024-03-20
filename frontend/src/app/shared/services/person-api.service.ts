/**
 * This service handles all interactions with the person API.
 * It can:
 *    - get all people for a department
 *    - get a specific person using department and person id
 *    - create a person
 *    - update a person
 *    - add department to a person
 *    - remove department from a person
 *    - add workplace (alias qualification) to a person
 *    - remove workplace from a person
 *    - add weekday (alias availability) to a person
 *    - remove weekday from a person
 *    - add absency to a person
 *    - get absency for a person by date
 *    - remove absency from a person
 */
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { constants } from '@app/core/constants/constants';
import { APIResponse } from '@app/core/interfaces/response';
import { Observable } from 'rxjs';
import { CreatePerson, PersonWithMetadata } from '../interfaces/person';
import { AbsenceReponse } from '@app/modules/absency/interfaces/absence';

@Injectable({
  providedIn: 'root',
})
export class PersonAPIService {
  private http = inject(HttpClient);

  getPerson(personID: string): Observable<APIResponse<PersonWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<PersonWithMetadata>>(url, httpOptions);
  }

  getPersons(departmentName: string): Observable<APIResponse<PersonWithMetadata[]>> {
    const url = `${constants.APIS.PLANNER}/person`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      params: {
        department: departmentName,
      },
      withCredentials: true,
    };

    return this.http.get<APIResponse<PersonWithMetadata[]>>(url, httpOptions);
  }

  createPerson(person: CreatePerson): Observable<APIResponse<PersonWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/person`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.post<APIResponse<PersonWithMetadata>>(url, JSON.stringify(person), httpOptions);
  }

  updatePerson(person: CreatePerson, id: string): Observable<APIResponse<PersonWithMetadata>> {
    const url = `${constants.APIS.PLANNER}/person/${id}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.put<APIResponse<PersonWithMetadata>>(url, person, httpOptions);
  }

  // Relations
  addDepartmentToPerson(departmentName: string, personID: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/department`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body = {
      department_id: departmentName,
    };

    return this.http.post<APIResponse<null>>(url, body, httpOptions);
  }

  removeDepartmentFromPerson(departmentName: string, personID: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/department/${departmentName}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }

  addWorkplaceToPerson(departmentName: string, workplaceName: string, personID: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/workplace`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body = {
      department_id: departmentName,
      workplace_id: workplaceName,
    };

    return this.http.post<APIResponse<null>>(url, body, httpOptions);
  }

  removeWorkplaceFromPerson(departmentName: string, workplaceName: string, personID: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/workplace`;

    const body = {
      department_id: departmentName,
      workplace_id: workplaceName,
    };

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      body: body,
    };

    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }

  addWeekdayToPerson(weekdayID: string, personID: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/weekday`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body = {
      weekday_name: weekdayID,
    };

    return this.http.post<APIResponse<null>>(url, body, httpOptions);
  }

  removeWeekdayFromPerson(weekdayID: string, personID: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/weekday/${weekdayID}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }

  addAbsencyToPerson(personID: string, date: string, reason: string | null): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/absency`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    const body = {
      date: date,
      reason: reason,
    };

    return this.http.post<APIResponse<null>>(url, body, httpOptions);
  }

  getAbsencyForPerson(personID: string, date: string): Observable<APIResponse<AbsenceReponse | null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/absency`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      params: new HttpParams({
        fromObject: {
          date: date,
        },
      }),
    };

    return this.http.get<APIResponse<AbsenceReponse | null>>(url, httpOptions);
  }

  getAbsencyForPersonInRange(personID: string, startDate: string, endDate: string): Observable<APIResponse<AbsenceReponse[]>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/absency`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      params: new HttpParams({
        fromObject: {
          start_date: startDate,
          end_date: endDate,
        },
      }),
    };

    return this.http.get<APIResponse<AbsenceReponse[]>>(url, httpOptions);
  }

  removeAbsencyFromPerson(personID: string, date: string): Observable<APIResponse<null>> {
    const url = `${constants.APIS.PLANNER}/person/${personID}/absency/${date}`;

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };
    return this.http.delete<APIResponse<null>>(url, httpOptions);
  }
}
