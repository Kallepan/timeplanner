import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSlideToggleHarness } from '@angular/material/slide-toggle/testing';
import { ActionsComponent } from './actions.component';
import { ActivatedRoute } from '@angular/router';

describe('ActionsComponent', () => {
  let component: ActionsComponent;
  let fixture: ComponentFixture<ActionsComponent>;
  let loader: HarnessLoader;
  let mockActivatedRoute: jasmine.SpyObj<ActivatedRoute>;

  beforeEach(async () => {
    mockActivatedRoute = jasmine.createSpyObj('ActivatedRoute', [''], {
      snapshot: {
        data: {
          title: 'Test',
          departmentId: 'departmentId',
        },
      },
    });

    await TestBed.configureTestingModule({
      imports: [ActionsComponent, MatSlideToggleModule],
      providers: [{ provide: ActivatedRoute, useValue: mockActivatedRoute }],
    }).compileComponents();

    fixture = TestBed.createComponent(ActionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should emit toggleColors', async () => {
    spyOn(component.toggleColors, 'emit');
    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleColors' }));
    await toggle.toggle();
    expect(component.toggleColors.emit).toHaveBeenCalledWith(false);

    await toggle.toggle();
    expect(component.toggleColors.emit).toHaveBeenCalledWith(true);
  });

  it('should emit toggleTimes', async () => {
    spyOn(component.toggleTimes, 'emit');

    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleTimes' }));
    await toggle.toggle();
    expect(component.toggleTimes.emit).toHaveBeenCalledWith(false);

    await toggle.toggle();
    expect(component.toggleTimes.emit).toHaveBeenCalledWith(true);
  });

  it('should display times', async () => {
    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleTimes' }));
    expect(await toggle.isChecked()).toBe(true);
  });

  it('should display colors', async () => {
    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleColors' }));
    expect(await toggle.isChecked()).toBe(true);
  });

  it('should emit shiftWeek positive', async () => {
    spyOn(component.shiftWeek, 'emit');
    const button = fixture.nativeElement.querySelector('#shift-forward-button');

    button.click();

    expect(component.shiftWeek.emit).toHaveBeenCalledWith(1);
  });

  it('should emit shiftWeek negative', async () => {
    spyOn(component.shiftWeek, 'emit');
    const button = fixture.nativeElement.querySelector('#shift-backward-button');

    button.click();

    expect(component.shiftWeek.emit).toHaveBeenCalledWith(-1);
  });

  it('should route to department', async () => {
    const editRouteButton = fixture.nativeElement.querySelector('#edit-route-button');
    expect(editRouteButton).toBeTruthy();
    expect(editRouteButton.textContent).toContain('Standardansicht');

    editRouteButton.click();
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      expect(location.pathname).toContain('planner');
      expect(location.pathname).toContain('departmentId');
    });
  });

  it('should emit toggleAbsencyPanel', async () => {
    spyOn(component.toggleAbsencyPanel, 'emit');
    const button = fixture.nativeElement.querySelector('#toggle-absency-panel-button');

    button.click();

    expect(component.toggleAbsencyPanel.emit).toHaveBeenCalled();
  });
});
