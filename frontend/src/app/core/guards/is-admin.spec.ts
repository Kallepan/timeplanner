import { NotificationService } from '../services/notification.service';
import { AuthService } from '../services/auth.service';
import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { isAdmin } from './auth-guards';
import { of } from 'rxjs';
import { UrlTree } from '@angular/router';

describe('isAdminGuard', () => {
  let mockAuthService: jasmine.SpyObj<AuthService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;

  beforeEach(async () => {
    mockAuthService = jasmine.createSpyObj('AuthService', ['isAdmin']);
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['warnMessage']);

    TestBed.configureTestingModule({
      imports: [RouterTestingModule],
      providers: [
        {
          provide: AuthService,
          useValue: mockAuthService,
        },
        {
          provide: NotificationService,
          useValue: mockNotificationService,
        },
      ],
    });
  });

  it('should be created', () => {
    expect(isAdmin).toBeTruthy();
  });

  it('should return URLTree if the user is not an admin', () => {
    mockAuthService.isAdmin.and.returnValue(of(false));

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const result = TestBed.runInInjectionContext(() => isAdmin());

    expect(mockAuthService.isAdmin).toHaveBeenCalled();
    result.subscribe((res) => {
      expect(mockNotificationService.warnMessage).toHaveBeenCalled();
      // url tree / is returned
      expect(res).toBeInstanceOf(UrlTree);
    });
  });

  it('should return true if the user is an admin', () => {
    mockAuthService.isAdmin.and.returnValue(of(true));

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const result = TestBed.runInInjectionContext(() => isAdmin());

    expect(mockAuthService.isAdmin).toHaveBeenCalled();
    result.subscribe((res) => {
      expect(mockNotificationService.warnMessage).not.toHaveBeenCalled();
      expect(res).toBeTrue();
    });
  });
});
