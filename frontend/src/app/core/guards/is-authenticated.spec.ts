import { NotificationService } from '../services/notification.service';
import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { isAuthenticated } from './auth-guard';
import { AuthService } from '../services/auth.service';

describe('isAuthenticatedGuard', () => {
  let mockAuthService: jasmine.SpyObj<AuthService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;

  beforeEach(async () => {
    mockAuthService = jasmine.createSpyObj('AuthService', ['isLoggedIn']);
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['warnMessage']);

    await TestBed.configureTestingModule({
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
    expect(isAuthenticated).toBeTruthy();
  });

  it('should return false if the user is not logged in', async () => {
    mockAuthService.isLoggedIn.and.returnValue(false);

    const result = await TestBed.runInInjectionContext(isAuthenticated);

    expect(mockNotificationService.warnMessage).toHaveBeenCalled();

    expect(result).toBeFalse();
  });

  it('should return true if the user is logged in', async () => {
    mockAuthService.isLoggedIn.and.returnValue(true);

    const result = await TestBed.runInInjectionContext(isAuthenticated);

    expect(result).toBeTrue();
  });
});
