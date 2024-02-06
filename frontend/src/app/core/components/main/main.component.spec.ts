import { TestBed, type ComponentFixture } from '@angular/core/testing';

import { OverlayContainer } from '@angular/cdk/overlay';
import { type HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { WritableSignal, signal } from '@angular/core';
import { MatSidenavHarness } from '@angular/material/sidenav/testing';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { MainComponent } from './main.component';

describe('MainComponent', () => {
  let mockOverlayContainer: jasmine.SpyObj<OverlayContainer>;
  let mockIsDarkSignal: WritableSignal<boolean>;
  let mockThemeHandlerService: jasmine.SpyObj<ThemeHandlerService>;

  let fixture: ComponentFixture<MainComponent>;
  let loader: HarnessLoader;

  beforeEach(() => {
    mockIsDarkSignal = signal(true);
    mockThemeHandlerService = jasmine.createSpyObj('ThemeHandlerService', ['toggleTheme'], {
      isDark$: mockIsDarkSignal(),
    });
    mockOverlayContainer = jasmine.createSpyObj('OverlayContainer', ['getContainerElement']);
    mockOverlayContainer.getContainerElement.and.returnValue(document.createElement('div'));

    TestBed.configureTestingModule({
      imports: [MainComponent, BrowserAnimationsModule, RouterTestingModule, HttpClientTestingModule, MatSnackBarModule],
      providers: [
        {
          provide: OverlayContainer,
          useValue: mockOverlayContainer,
        },
        { provide: ThemeHandlerService, useValue: mockThemeHandlerService },
      ],
    });
    fixture = TestBed.createComponent(MainComponent);
    fixture.detectChanges();
    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(MainComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  // Test for theme
  it('should have a default theme mode', () => {
    const fixture = TestBed.createComponent(MainComponent);
    const app = fixture.componentInstance;
    expect(app.themeMode).toEqual('theme-dark');
  });

  // Toggle theme test
  it('should toggle theme mode', () => {
    const fixture = TestBed.createComponent(MainComponent);
    const app = fixture.componentInstance;
    app.toggleTheme();

    expect(mockThemeHandlerService.toggleTheme).toHaveBeenCalled();
  });

  // Sidenav testing
  it('should open and close sidenav', async () => {
    // Fetch the sidenav element
    const sidenav = await loader.getHarness(MatSidenavHarness);

    // Check if the sidenav is closed by default
    expect(await sidenav.isOpen()).toBe(false);
  });

});

describe('MainComponent (injected)', () => {
  let mockOverlayContainer: jasmine.SpyObj<OverlayContainer>;
  let themeHandlerService: ThemeHandlerService;

  let fixture: ComponentFixture<MainComponent>;
  let loader: HarnessLoader;

  beforeEach(() => {
    mockOverlayContainer = jasmine.createSpyObj('OverlayContainer', ['getContainerElement']);
    mockOverlayContainer.getContainerElement.and.returnValue(document.createElement('div'));

    TestBed.configureTestingModule({
      imports: [MainComponent, BrowserAnimationsModule, RouterTestingModule, HttpClientTestingModule, MatSnackBarModule],
      providers: [
        {
          provide: OverlayContainer,
          useValue: mockOverlayContainer,
        },
        { provide: ThemeHandlerService },
      ],
    });
    fixture = TestBed.createComponent(MainComponent);
    themeHandlerService = TestBed.inject(ThemeHandlerService);
    fixture.detectChanges();
    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should call getContainerElement of OverlayContainer', () => {
    // call on init
    expect(mockOverlayContainer.getContainerElement).toHaveBeenCalledTimes(2);

    fixture.componentInstance.toggleTheme();
    fixture.detectChanges();

    // The method should be called again when the theme changes
    expect(mockOverlayContainer.getContainerElement).toHaveBeenCalledTimes(4);

    fixture.componentInstance.toggleTheme();
    fixture.detectChanges();

    // The method should be called again when the theme changes
    expect(mockOverlayContainer.getContainerElement).toHaveBeenCalledTimes(6);
  });
});
