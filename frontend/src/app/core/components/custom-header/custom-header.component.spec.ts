import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { NotificationService } from '@app/core/services/notification.service';
import { of } from 'rxjs';
import { CustomHeaderComponent } from './custom-header.component';

describe('CustomHeaderComponent', () => {
  let component: CustomHeaderComponent;
  let fixture: ComponentFixture<CustomHeaderComponent>;
  let notificationService: jasmine.SpyObj<NotificationService>;
  let router: jasmine.SpyObj<Router>;
  let activatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(() => {
    notificationService = jasmine.createSpyObj('NotificationService', [
      'infoMessage',
      'warnMessage',
    ]);
    router = jasmine.createSpyObj('Router', ['navigate'], {
      events: of(new NavigationEnd(1, '', '')),
    });
    activatedRoute = jasmine.createSpyObj('ActivatedRoute', ['data'], {
      data: of({ featureFlag: 'BAK' }),
    });
  });

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CustomHeaderComponent],
      providers: [
        { provide: NotificationService, useValue: notificationService },
        { provide: Router, useValue: router },
        { provide: ActivatedRoute, useValue: activatedRoute },
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(CustomHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
