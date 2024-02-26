import { RouterTestingModule } from '@angular/router/testing';
import { AuthService } from '../services/auth.service';
import { NotificationService } from '../services/notification.service';
import { TestBed } from '@angular/core/testing';
import { ActivatedRouteSnapshot, UrlTree } from '@angular/router';
import { of } from 'rxjs';
import { hasAccessToDepartmentGuard } from './auth-guards';

describe('hasAccessToDepartmentGuard', () => {
  let mockAuthService: jasmine.SpyObj<AuthService>;
  let mockNotificationService: jasmine.SpyObj<NotificationService>;
  let route: ActivatedRouteSnapshot;

  beforeEach(async () => {
    mockAuthService = jasmine.createSpyObj('AuthService', ['hasAccessToDepartment']);
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

    route = new ActivatedRouteSnapshot();
    route.queryParams = { department: 'testDepartment' };
  });

  it('should be created', () => {
    expect(mockAuthService).toBeTruthy();
  });

  it('should return true if the user has access to the department', async () => {
    mockAuthService.hasAccessToDepartment.and.returnValue(of(true));

    route.queryParams = { department: 'testDepartment' };

    const res = TestBed.runInInjectionContext(() => hasAccessToDepartmentGuard(route));

    expect(mockAuthService.hasAccessToDepartment).toHaveBeenCalledWith('testDepartment');
    res.subscribe((result) => {
      expect(result).toBeTrue();
    });
  });

  it('should return URLTree if the user does not have access to the department', async () => {
    mockAuthService.hasAccessToDepartment.and.returnValue(of(false));

    route.queryParams = { department: 'testDepartment' };

    const res = TestBed.runInInjectionContext(() => hasAccessToDepartmentGuard(route));

    expect(mockAuthService.hasAccessToDepartment).toHaveBeenCalledWith('testDepartment');
    res.subscribe((result) => {
      expect(mockNotificationService.warnMessage).toHaveBeenCalled();
      expect(result).toBeInstanceOf(UrlTree);
    });
  });
});
