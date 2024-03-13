import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeComponent } from './home.component';
import { RouterTestingModule } from '@angular/router/testing';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { AuthService } from '@app/core/services/auth.service';

describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;
  let mockAuthService: jasmine.SpyObj<AuthService>;

  beforeEach(() => {
    mockAuthService = jasmine.createSpyObj('AuthService', ['isAdmin']);

    TestBed.configureTestingModule({
      imports: [HomeComponent, RouterTestingModule, MatButtonModule, MatTooltipModule],
      providers: [
        {
          provide: AuthService,
          useValue: mockAuthService,
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
