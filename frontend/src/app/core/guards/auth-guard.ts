import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { map } from 'rxjs';

// Guard to check if user has access to a department, TODO: fix me
export const hasAccessToDepartmentGuard: CanActivateFn = (route: ActivatedRouteSnapshot) => {
  const authService = inject(AuthService);
  const router = inject(Router);

  // get department by query param
  const departmentFlag = route.queryParams['department'];

  // simply catch the error and return false if the user does not have access
  return authService.hasAccessToDepartment(departmentFlag).pipe(map((hasAccess) => hasAccess || router.createUrlTree(['/'])));
};

// Simple guard to check if the user is authenticated
export const isAuthenticated = () => inject(AuthService).isLoggedIn();
