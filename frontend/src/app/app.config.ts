import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { ApplicationConfig, importProvidersFrom } from '@angular/core';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideRouter } from '@angular/router';
import { httpErrorInterceptor } from '@core/interceptors/http-error-interceptor';
import { NotificationService } from '@core/services/notification.service';
import { routes } from './app.routes';
import { MAT_DATE_LOCALE, provideNativeDateAdapter } from '@angular/material/core';
import { timeoutInterceptor } from './core/interceptors/timeout-interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideAnimations(),
    provideHttpClient(withInterceptors([httpErrorInterceptor, timeoutInterceptor])),
    importProvidersFrom(MatSnackBarModule),
    NotificationService,
    provideNativeDateAdapter(),
    {
      provide: MAT_DATE_LOCALE,
      useValue: 'de-DE',
    },
  ],
};
