import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable, computed, inject, signal, type WritableSignal } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, catchError, from, map, of, switchMap, tap } from 'rxjs';
import { constants } from '../../constants/constants';
import { messages } from '../../constants/messages';
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

  initialized = computed(() => {
    return this._authData() !== undefined;
  });

  isLoggedIn = computed(() => {
    // Note this property gets populated by verifyLogin() at ngOnInit() in app.component.ts
    return this._authData() !== null;
  });

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

  verifyLogin(): void {
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
        catchError(() => of(null)),
      )
      .subscribe({
        next: (data) => {
          this._authData.set(data);
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
      .pipe(map((resp) => resp.data))
      .subscribe({
        next: (data) => {
          this._loading.set(false);
          this._authData.set(data);
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
