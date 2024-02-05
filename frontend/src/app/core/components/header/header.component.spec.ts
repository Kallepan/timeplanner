import { TestBed, type ComponentFixture } from '@angular/core/testing';

import { type HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { MatSlideToggleHarness } from '@angular/material/slide-toggle/testing';
import { By } from '@angular/platform-browser';
import { HeaderComponent } from './header.component';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;
  let loader: HarnessLoader;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HeaderComponent],
      providers: [provideHttpClientTesting(), provideHttpClient()],
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

  it('should emit toggle theme event', async () => {
    spyOn(component.themeToggled, 'emit');

    const matSlideToggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleTheme' }));
    expect(await matSlideToggle.isChecked()).toBe(false);

    await matSlideToggle.toggle();

    expect(component.themeToggled.emit).toHaveBeenCalled();
  });
});
