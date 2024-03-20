import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeComponent } from './home.component';
import { AuthService } from '@app/core/services/auth.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';

describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;
  let mockAuthService: jasmine.SpyObj<AuthService>;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(() => {
    mockAuthService = jasmine.createSpyObj('AuthService', ['isAdmin']);
    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['activeDepartment']);
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [], {
      queryParams: of({ department: 'test' }),
    });

    TestBed.configureTestingModule({
      imports: [HomeComponent],
      providers: [
        {
          provide: AuthService,
          useValue: mockAuthService,
        },
        {
          provide: ActiveDepartmentHandlerService,
          useValue: mockActiveDepartmentHandlerService,
        },
        {
          provide: ActivatedRoute,
          useValue: mockActivatedRoute,
        },
      ],
    });
    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have as many buttons as links', () => {
    // Get number of buttons
    const compiled = fixture.debugElement.nativeElement;
    const buttons = compiled.querySelectorAll('button');

    // Get number of links
    expect(component.routeConfigurations.length).toEqual(buttons.length);
  });

  it('should have an admin button if the user is an admin', () => {
    spyOn(component, 'isAdmin').and.returnValue(true);
    fixture.detectChanges();

    const compiled = fixture.debugElement.nativeElement;
    const buttons = compiled.querySelectorAll('#admin-route-button');
    const adminButton = buttons[buttons.length - 1];

    expect(adminButton).toBeTruthy();
    expect(adminButton.textContent).toContain('Admin');
  });

  it('should not have an admin button if the user is not an admin', () => {
    spyOn(component, 'isAdmin').and.returnValue(false);
    fixture.detectChanges();

    const compiled = fixture.debugElement.nativeElement;
    const buttons = compiled.querySelectorAll('#admin-route-button');

    expect(buttons.length).toEqual(0);
  });
});
