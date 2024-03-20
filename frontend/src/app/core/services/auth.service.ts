import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable, computed, inject, signal, type WritableSignal } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, catchError, from, map, of, switchMap, tap, throwError } from 'rxjs';
import { constants } from '../constants/constants';
import { messages } from '../constants/messages';
import { NotificationService } from './notification.service';
import { APIResponse } from '../interfaces/response';

type AuthResponse = {
  username: string;
  email: string;
};

type AuthData = {
  username: string;
  email: string;
};

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly http = inject(HttpClient);
  private readonly _notificationService = inject(NotificationService);
  private readonly _router = inject(Router);

  private readonly _authData = signal<AuthData | null | undefined>(undefined);

  private _loading = signal(false);
  get loading$() {
    return this._loading();
  }

  private _isAdmin = signal(false);
  get isAdmin$() {
    return this._isAdmin();
  }

  initialized = computed(() => {
    return this._authData() !== undefined;
  });

  isLoggedIn = computed(() => {
    // Note this property gets populated by verifyToken() at ngOnInit() in app.component.ts
    return this._authData() !== null;
  });

  isAdmin(): Observable<boolean> {
    // Check if the user is an admin used to live check if the user is admin to acess routes
    // alternatively used during login or verifyToken() to check if the user is an admin
    // and store the result in a signal (this._isAdmin) to be used in the app.
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    return this.http.get<APIResponse<boolean | null>>(`${constants.APIS.AUTH}/check-admin`, httpOptions).pipe(
      map((resp) => resp.data),
      map((isAdmin) => isAdmin === true),
      catchError(() => of(false)),
      tap((isAdmin) => this._isAdmin.set(isAdmin)),
    );
  }

  hasAccessToDepartment(departmentName: string): Observable<boolean> {
    // Check if the user has access to a department
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
      params: new HttpParams({
        fromObject: {
          department: departmentName,
        },
      }),
    };

    return this.http.get<null>(`${constants.APIS.AUTH}/me`, httpOptions).pipe(
      map(() => true),
      catchError(() => of(false)),
    );
  }

  verifyToken(): void {
    /* Called at ngOnInit() in app.component.ts to check if the user is logged in using cookies */
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    this.http
      .get<APIResponse<AuthResponse>>(`${constants.APIS.AUTH}/me`, httpOptions)
      .pipe(
        map((resp) => resp.data),
        tap((data) => this._authData.set(data)),
        switchMap(() => this.isAdmin()),
        tap((data) => this._isAdmin.set(data)),
      )
      .subscribe({
        next: () => {
          this._loading.set(false);
          this._notificationService.infoMessage(messages.AUTH.LOGGED_IN);
        },
        error: () => {
          this._loading.set(false);
          this._authData.set(null);
        },
      });
  }

  login(username: string | null, password: string | null): void {
    const data = {
      username,
      password,
    };

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };
    this._loading.set(true);
    this.http
      .post<APIResponse<AuthResponse>>(`${constants.APIS.AUTH}/login`, data, httpOptions)
      .pipe(
        map((resp) => resp.data),
        tap((data) => this._authData.set(data)),
        switchMap(() => this.isAdmin()),
        tap((data) => this._isAdmin.set(data)),
        catchError((err) => {
          this._authData.set(null);
          return throwError(() => err);
        }),
      )
      .subscribe({
        next: () => {
          this._loading.set(false);
          this._notificationService.infoMessage(messages.AUTH.LOGGED_IN);
        },
        error: () => {
          this._loading.set(false);
          this._notificationService.warnMessage(messages.AUTH.LOGIN_FAILED);
        },
      });
  }

  logout(): void {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      withCredentials: true,
    };

    this.http
      .post<null>(`${constants.APIS.AUTH}/logout`, {}, httpOptions)
      .pipe(
        tap(() => {
          this._authData.set(null);
        }),
        switchMap(() => from(this._router.navigate(['']))),
      )
      .subscribe({
        next: () => {
          this._notificationService.infoMessage(messages.AUTH.LOGGED_OUT);
        },
      });
  }

  get authData(): WritableSignal<AuthData | null | undefined> {
    return this._authData;
  }
}
