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

  it('should emit toggleComments', async () => {
    spyOn(component.toggleComments, 'emit');
    const toggle = await loader.getHarness(MatSlideToggleHarness.with({ selector: '#toggleComments' }));
    await toggle.toggle();
    expect(component.toggleComments.emit).toHaveBeenCalledWith(false);

    await toggle.toggle();
    expect(component.toggleComments.emit).toHaveBeenCalledWith(true);
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
    const routerLink = fixture.nativeElement.querySelector('#edit-route-button');
    expect(routerLink.getAttribute('href')).toBe('/department/departmentId');
  });
});
