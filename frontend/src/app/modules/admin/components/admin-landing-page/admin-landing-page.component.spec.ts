import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminLandingPageComponent } from './admin-landing-page.component';
import { ActivatedRoute } from '@angular/router';

describe('AdminLandingPageComponent', () => {
  let component: AdminLandingPageComponent;
  let fixture: ComponentFixture<AdminLandingPageComponent>;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', ['snapshot'], { snapshot: { data: { title: 'Test' } } });

    await TestBed.configureTestingModule({
      imports: [AdminLandingPageComponent],
      providers: [{ provide: ActivatedRoute, useValue: mockActivatedRoute }],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have a list of routes', () => {
    expect(component.routes).toBeDefined();
  });

  it('should have must have routes', () => {
    expect(component.routes).toContain({ path: 'schema', name: 'Schemaeditor' });
    expect(component.routes).toContain({ path: 'person', name: 'Personaleditor' });
    expect(component.routes).toContain({ path: 'workday', name: 'Arbeitstagseditor' });
  });

  it('should display all routes in the list', () => {
    const compiled = fixture.nativeElement;
    const links = compiled.querySelectorAll('button');
    expect(links.length).toBe(component.routes.length);
  });
});
