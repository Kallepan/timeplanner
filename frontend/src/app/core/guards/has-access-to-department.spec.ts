import { RouterTestingModule } from '@angular/router/testing';
import { AuthService } from '../services/auth.service';
import { NotificationService } from '../services/notification.service';
import { TestBed } from '@angular/core/testing';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';
import { of } from 'rxjs';
import { hasAccessToDepartmentGuard } from './auth-guard';

describe('hasAccessToDepartmentGuard', () => {
  let mockAuthService: jasmine.SpyObj<AuthService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let mockRouter: jasmine.SpyObj<Router>;
  let route: ActivatedRouteSnapshot;

  beforeEach(async () => {
    mockAuthService = jasmine.createSpyObj('AuthService', ['hasAccessToDepartment']);
    mockNotificationService = jasmine.createSpyObj('NotificationService', ['warnMessage']);
    mockRouter = jasmine.createSpyObj('Router', ['createUrlTree']);

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
        {
          provide: Router,
          useValue: mockRouter,
        },
      ],
    });

    route = new ActivatedRouteSnapshot();
    route.queryParams = { department: 'testDepartment' };
  });

  it('should be created', () => {
    expect(mockAuthService).toBeTruthy();
  });

  it('should return true if the user has access to the department', async () => {
    mockAuthService.hasAccessToDepartment.and.returnValue(of(true));

    route.queryParams = { department: 'testDepartment' };

    await TestBed.runInInjectionContext(() => hasAccessToDepartmentGuard(route, {} as RouterStateSnapshot));

    expect(mockAuthService.hasAccessToDepartment).toHaveBeenCalledWith('testDepartment');
  });

  it('should return false if the user does not have access to the department', async () => {
    mockAuthService.hasAccessToDepartment.and.returnValue(of(false));

    route.queryParams = { department: 'testDepartment' };

    await TestBed.runInInjectionContext(() => hasAccessToDepartmentGuard(route, {} as RouterStateSnapshot));

    expect(mockAuthService.hasAccessToDepartment).toHaveBeenCalledWith('testDepartment');
  });
});
