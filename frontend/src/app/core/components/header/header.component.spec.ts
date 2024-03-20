/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @angular-eslint/directive-selector */
/* eslint-disable @angular-eslint/directive-class-suffix */
import { TestBed, fakeAsync, type ComponentFixture } from '@angular/core/testing';

import { type HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { MatSlideToggleHarness } from '@angular/material/slide-toggle/testing';
import { By } from '@angular/platform-browser';
import { HeaderComponent } from './header.component';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { constants } from '@app/core/constants/constants';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;
  let loader: HarnessLoader;
  let mockActivatedRotue: jasmine.SpyObj<ActivatedRoute>;
  let routerLinks: any[];
  let linkDes: any[];

  beforeEach(() => {
    mockActivatedRotue = jasmine.createSpyObj('ActivatedRoute', ['snapshot'], { snapshot: { data: { title: 'Test' } } });

    TestBed.configureTestingModule({
      imports: [HeaderComponent],
      providers: [provideHttpClientTesting(), provideHttpClient(), { provide: ActivatedRoute, useValue: mockActivatedRotue }],
    });
    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('title should be visible', () => {
    const node = fixture.debugElement.query(By.css('.header-title'));
    expect(node.nativeElement.textContent).toContain(component.title);
  });

  it('should emit open sidenav event', () => {
    spyOn(component.sidenavToggled, 'emit');
    // Fetch by 'button' tag
    const nodes = fixture.debugElement.nativeElement.querySelectorAll('button');
    nodes[0].click();

    expect(component.sidenavToggled.emit).toHaveBeenCalled();
  });

  it('should emit toggle theme event', fakeAsync(async () => {
    spyOn(component.themeToggled, 'emit');

    const matSlideToggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleTheme' }));
    expect(await matSlideToggle.isChecked()).toBe(false);

    await matSlideToggle.toggle();

    expect(component.themeToggled.emit).toHaveBeenCalled();
  }));

  it('should have correct title', () => {
    expect(component.title).toBe(constants.TITLE_SHORT);
  });

  it('should fetch routerLinks', () => {
    linkDes = fixture.debugElement.queryAll(By.directive(RouterLink));
    routerLinks = linkDes.map((de) => de.injector.get(RouterLink));

    expect(routerLinks.length).toBe(1);
    expect(routerLinks[0].navigatedTo).toBe(undefined);
    expect(routerLinks[0].commands[0]).toBe('/');
  });
});
