import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateAbsencyDialogComponent, CreateAbsencyDialogData } from './create-absency-dialog.component';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { provideNativeDateAdapter } from '@angular/material/core';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatRadioButtonHarness } from '@angular/material/radio/testing';
import { MatDatepickerInputHarness } from '@angular/material/datepicker/testing';
import { formatDateToDateString } from '@app/shared/functions/format-date-to-string.function';

describe('CreateAbsencyDialogComponent', () => {
  let component: CreateAbsencyDialogComponent;
  let fixture: ComponentFixture<CreateAbsencyDialogComponent>;
  let loader: HarnessLoader;

  let mockDialogData: CreateAbsencyDialogData;

  beforeEach(async () => {
    mockDialogData = {
      personID: 'Mock Person ID',
      startDate: new Date(),
    };

    await TestBed.configureTestingModule({
      imports: [CreateAbsencyDialogComponent],
      providers: [{ provide: MAT_DIALOG_DATA, useValue: mockDialogData }, provideNativeDateAdapter(), provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAbsencyDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize control with start date', () => {
    expect(component.control.get('endDate')?.value).toEqual(mockDialogData.startDate);
  });

  it('should initialize title with person ID', () => {
    expect(component.title).toEqual(`Abwesenheit erstellen für ${mockDialogData.personID}`);
  });

  it('should initialize control with empty reason', () => {
    expect(component.control.get('reason')?.value).toEqual('');
  });

  it('should initialize absency reasons', async () => {
    expect(component.absencyReasons).toBeDefined();

    // get radio buttons
    const radioButtons = await loader.getAllHarnesses(MatRadioButtonHarness);

    expect(radioButtons).toBeDefined();
    expect(radioButtons.length).toEqual(component.absencyReasons.length);

    // compare radio button labels with absency reasons
    radioButtons.forEach(async (radioButton, index) => {
      const label = await radioButton.getLabelText();
      expect(label).toEqual(component.absencyReasons[index]);
    });
  });

  it('should set reason', async () => {
    const radioButtons = await loader.getAllHarnesses(MatRadioButtonHarness);

    // select first radio button
    await radioButtons[0].check();

    expect(component.control.get('reason')?.value).toEqual(component.absencyReasons[0]);
  });

  it('should open date picker', async () => {
    const datePicker = await loader.getHarness(MatDatepickerInputHarness);

    expect(datePicker).toBeDefined();

    // open date picker
    await datePicker.openCalendar();
    expect(await datePicker.isCalendarOpen()).toBeTruthy();

    // dates after today should be enabled
    datePicker.getMin().then((minDate) => {
      const today = new Date();
      const formattedDate = formatDateToDateString(today);
      expect(minDate).toEqual(formattedDate);
    });

    // close date picker
    await datePicker.closeCalendar();
    expect(await datePicker.isCalendarOpen()).toBeFalsy();
  });

  it('form should be invalid by default', () => {
    expect(component.control.valid).toBeFalsy();
  });
});
