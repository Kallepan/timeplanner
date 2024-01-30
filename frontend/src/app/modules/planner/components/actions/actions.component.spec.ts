import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ActionsComponent } from './actions.component';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatSlideToggleHarness } from '@angular/material/slide-toggle/testing';

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

  it('should emit toggleTimeLabel', async () => {
    spyOn(component.toggleTimeLabel, 'emit');

    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleTimeLabel' }));
    await toggle.toggle();
    expect(component.toggleTimeLabel.emit).toHaveBeenCalledWith(false);

    await toggle.toggle();
    expect(component.toggleTimeLabel.emit).toHaveBeenCalledWith(true);
  });

  it('should emit togglePersonsLabel', async () => {
    spyOn(component.togglePersonsLabel, 'emit');
    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#togglePersonsLabel' }));
    await toggle.toggle();
    expect(component.togglePersonsLabel.emit).toHaveBeenCalledWith(false);

    await toggle.toggle();
    expect(component.togglePersonsLabel.emit).toHaveBeenCalledWith(true);
  });

  it('should emit toggleEditingMode', async () => {
    spyOn(component.toggleEditingMode, 'emit');
    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleEditingMode' }));
    await toggle.toggle();
    expect(component.toggleEditingMode.emit).toHaveBeenCalledWith(false);

    await toggle.toggle();
    expect(component.toggleEditingMode.emit).toHaveBeenCalledWith(true);
  });
});
