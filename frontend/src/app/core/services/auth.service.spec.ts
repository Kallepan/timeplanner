import { TestBed } from '@angular/core/testing';

import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { constants } from '../../constants/constants';
import { AuthService } from './auth.service';
import { NotificationService } from './notification.service';
import { messages } from '@app/constants/messages';

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;
  let notificationService: jasmine.SpyObj<NotificationService>;

  beforeEach(() => {
    notificationService = jasmine.createSpyObj('NotificationService', ['infoMessage', 'warnMessage']);

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

  it('login failed', () => {
    // login failed
    service.login('test', 'test');
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/login`);
    expect(req.request.method).toBe('POST');
    req.flush({}, { status: 401, statusText: 'Unauthorized' });

    // should call notificationService.warnMessage
    expect(notificationService.warnMessage).toHaveBeenCalledWith(messages.AUTH.LOGIN_FAILED);
  });

  it('login should set authData', () => {
    service.login('test', 'test');
    const loginReq = httpMock.expectOne(`${constants.APIS.AUTH}/login`);

    loginReq.flush({ data: { username: 'test', email: 'test@example.com' } });
    expect(loginReq.request.method).toBe('POST');
    expect(service.authData()).toEqual({ username: 'test', email: 'test@example.com' });
    const adminReq = httpMock.expectOne(`${constants.APIS.AUTH}/check-admin`);

    adminReq.flush({ data: true });
    expect(adminReq.request.method).toBe('GET');

    // should call notificationService.infoMessage
    expect(notificationService.infoMessage).toHaveBeenCalledWith(messages.AUTH.LOGGED_IN);
    expect(service.isAdmin$).toBeTrue();
  });

  it('isLoggedIn should return true', () => {
    service.login('test', 'test');
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/login`);
    expect(req.request.method).toBe('POST');
    req.flush({ data: { username: 'test', email: 'test@example.com' } });
    expect(service.isLoggedIn()).toBeTrue();
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

  it('should display error message if login failed', () => {
    service.login('test', 'test');
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/login`);
    expect(req.request.method).toBe('POST');
    req.flush({}, { status: 401, statusText: 'Unauthorized' });
    expect(notificationService.warnMessage).toHaveBeenCalledWith(messages.AUTH.LOGIN_FAILED);
  });

  it('isAdmin should return true if server returns true', () => {
    service.isAdmin().subscribe((result) => expect(result).toBeTrue());
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/check-admin`);
    expect(req.request.method).toBe('GET');
    req.flush({ data: true, status: 200, statusText: 'OK' });
  });

  it('isAdmin should return false if server returns false', () => {
    service.isAdmin().subscribe((result) => expect(result).toBeFalse());
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/check-admin`);
    expect(req.request.method).toBe('GET');
    req.flush(
      {
        data: false,
      },
      { status: 401, statusText: 'Unauthorized' },
    );
  });

  it('verifyToken should call correct methods', () => {
    service.verifyToken();
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/me`);
    expect(req.request.method).toBe('GET');
    req.flush({ data: { username: 'test', email: 'test@example.com' } });

    const adminReq = httpMock.expectOne(`${constants.APIS.AUTH}/check-admin`);
    adminReq.flush({ data: true });

    expect(service.authData()).toEqual({ username: 'test', email: 'test@example.com' });
    expect(service.isAdmin$).toBeTrue();
    expect(notificationService.infoMessage).toHaveBeenCalledWith(messages.AUTH.LOGGED_IN);
    expect(service.loading$).toBeFalse();
  });

  it('verifyToken should call correct methods and display error message', () => {
    service.verifyToken();
    const req = httpMock.expectOne(`${constants.APIS.AUTH}/me`);
    expect(req.request.method).toBe('GET');
    req.flush({}, { status: 401, statusText: 'Unauthorized' });

    expect(service.authData()).toBeNull();
    expect(notificationService.infoMessage).not.toHaveBeenCalled();
    expect(service.isAdmin$).toBeFalse();
    expect(service.loading$).toBeFalse();
  });
});
