import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { map, tap } from 'rxjs';
import { messages } from '@app/constants/messages';
import { NotificationService } from '../services/notification.service';

// Guard to check if user has access to a department
export const hasAccessToDepartmentGuard: CanActivateFn = (route: ActivatedRouteSnapshot) => {
  const authService = inject(AuthService);
  const router = inject(Router);
  const notificationService = inject(NotificationService);

  // get department by query param
  const departmentFlag = route.queryParams['department'];

  // simply catch the error and return false if the user does not have access
  return authService.hasAccessToDepartment(departmentFlag).pipe(
    tap((hasAccess) => {
      if (!hasAccess) notificationService.warnMessage(messages.AUTH.FORBIDDEN);
    }),
    map((hasAccess) => hasAccess || router.createUrlTree(['/'])),
  );
};

// Simple guard to check if the user is authenticated
export const isAuthenticated = () => {
  const notificationService = inject(NotificationService);
  const isLoggedIn = inject(AuthService).isLoggedIn();

  if (!isLoggedIn) notificationService.warnMessage(messages.AUTH.UNAUTHORIZED);

  return isLoggedIn;
};
