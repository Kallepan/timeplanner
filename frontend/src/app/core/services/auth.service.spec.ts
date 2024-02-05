import { TestBed } from '@angular/core/testing';

import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { constants } from '../../constants/constants';
import { AuthService } from './auth.service';
import { NotificationService } from './notification.service';

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;
  let notificationService: jasmine.SpyObj<NotificationService>;

  beforeEach(() => {
    notificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'errorMessage']);

    TestBed.configureTestingModule({
      imports: [MatSnackBarModule, BrowserAnimationsModule],
      providers: [{ provide: NotificationService, useValue: notificationService }, provideHttpClient(), provideHttpClientTesting()],
    });
    service = TestBed.inject(AuthService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('authData should return null', () => {
    expect(service.authData()).toBeUndefined();
  });

  it('logout should set authData to null', () => {
    service.logout();
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/logout`);
    expect(req.request.method).toBe('POST');
    req.flush(null);
    expect(service.authData()).toBeNull();
  });

  it('login should set authData', () => {
    service.login('test', 'test');
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/login`);
    expect(req.request.method).toBe('POST');
    req.flush({data: { username: 'test', email: 'test@example.com' }});
    expect(service.authData()).toEqual({ username: 'test', email: 'test@example.com' });
  });

  it('hasAccessToDepartment should return true', () => {
    service.hasAccessToDepartment('test').subscribe((result) => expect(result).toBeTrue());
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/me?department=test`);
    expect(req.request.method).toBe('GET');
    req.flush({});
  });

  it('hasAccessToDepartment should return false', () => {
    service.hasAccessToDepartment('test').subscribe((result) => expect(result).toBeFalse());
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/me?department=test`);
    expect(req.request.method).toBe('GET');
    req.flush({}, { status: 401, statusText: 'Unauthorized' });
  });
});
