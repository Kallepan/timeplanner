import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AbsencyLandingPageComponent } from './absency-landing-page.component';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { provideNoopAnimations } from '@angular/platform-browser/animations';

describe('AbsencyLandingPageComponent', () => {
  let component: AbsencyLandingPageComponent;
  let fixture: ComponentFixture<AbsencyLandingPageComponent>;

  // mocks
  let mockDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;
  let mockActivePersonHandlerService: jasmine.SpyObj<ActivePersonHandlerServiceService>;
  let mockPersonDataContainerService: jasmine.SpyObj<PersonDataContainerService>;

  beforeEach(async () => {
    mockDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', ['']);
    mockActivePersonHandlerService = jasmine.createSpyObj('ActivePersonHandlerServiceService', ['']);
    mockPersonDataContainerService = jasmine.createSpyObj('PersonDataContainerService', ['']);

    await TestBed.configureTestingModule({
      providers: [
        { provide: ActiveDepartmentHandlerService, useValue: mockDepartmentHandlerService },
        { provide: ActivePersonHandlerServiceService, useValue: mockActivePersonHandlerService },
        { provide: PersonDataContainerService, useValue: mockPersonDataContainerService },
        provideNoopAnimations(),
      ],
      imports: [AbsencyLandingPageComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AbsencyLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
