import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSlideToggleHarness } from '@angular/material/slide-toggle/testing';
import { ActionsComponent } from './actions.component';

describe('ActionsComponent', () => {
  let component: ActionsComponent;
  let fixture: ComponentFixture<ActionsComponent>;
  let loader: HarnessLoader;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ActionsComponent, MatSlideToggleModule],
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
});
