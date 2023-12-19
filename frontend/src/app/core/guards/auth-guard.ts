import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivateFn,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { constants } from '@app/constants/constants';
import { catchError, map, of } from 'rxjs';
import { AuthService } from '../services/auth.service';

// Guard to check if a feature flag is enabled for the user
export const featureFlagGuard: CanActivateFn = (
  route: ActivatedRouteSnapshot,
  state: RouterStateSnapshot,
) => {
  const router = inject(Router);
  const http = inject(HttpClient);

  const httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
    }),
    withCredentials: true,
  };

  const requestedFeatureFlag = (route.data['featureFlag'] as string)
    .trim()
    .toUpperCase();

  return http
    .get<any>(
      `${constants.APIS.AUTH}/has_access/${requestedFeatureFlag}`,
      httpOptions,
    )
    .pipe(
      map(() => true),
      catchError(() => of(false)),
      // If the feature flag is disabled, redirect the user to the home page
      map((isEnabled) => isEnabled || router.createUrlTree([''])),
    );
};

// Simple guard to check if the user is authenticated
export const isAuthenticated = () => inject(AuthService).isLoggedIn();
