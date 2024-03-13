import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditWeekdaysOfTimeslotComponent } from './edit-weekdays-of-timeslot.component';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { FormControl } from '@angular/forms';

describe('EditWeekdaysOfTimeslotComponent', () => {
  let component: EditWeekdaysOfTimeslotComponent;
  let fixture: ComponentFixture<EditWeekdaysOfTimeslotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EditWeekdaysOfTimeslotComponent],
      providers: [provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(EditWeekdaysOfTimeslotComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should not have a timeslot at initialization', () => {
    expect(component.timeslot).toBeNull();
  });

  it('should display timeslot name', () => {
    component.timeslot = {
      id: '1',
      name: 'test',
      active: true,
      department_id: 'test',
      workplace_id: 'test',
      weekdays: [
        {
          id: 1,
          name: 'test',
          start_time: 'test',
          end_time: 'test',
        },
      ],

      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    };

    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h3').textContent).toContain(component.timeslot.name);
  });

  it('submitEdit should emit editRequest event', () => {
    spyOn(component.editRequest, 'emit');
    component.formGroup.controls['createForm'].setValue({
      weekday: 1,
      startTime: 'test',
      endTime: 'test',
    });
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    component.timeslot = {} as any;
    component.submitEdit(new FormControl(), new FormControl(), 1);
    expect(component.editRequest.emit).toHaveBeenCalled();
  });

  it('submitEdit should not emit editRequest if timeslot is undefined', () => {
    spyOn(component.editRequest, 'emit');
    component.formGroup.controls['createForm'].setValue({
      weekday: 1,
      startTime: 'test',
      endTime: 'test',
    });
    component.submitEdit(new FormControl(), new FormControl(), 1);
    expect(component.editRequest.emit).not.toHaveBeenCalled();
  });

  it('submitEdit should not emit editRequest if controlform is undefined', () => {
    spyOn(component.editRequest, 'emit');
    component.formGroup.controls['createForm'].setValue({
      weekday: 1,
      startTime: 'test',
      endTime: 'test',
    });
    const dummyFormControl = new FormControl();
    dummyFormControl.setErrors({ pattern: true });

    component.submitEdit(new FormControl(), dummyFormControl, 1);
    expect(component.editRequest.emit).not.toHaveBeenCalled();
  });

  it('submitRemove should emit event', () => {
    spyOn(component.removeRequest, 'emit');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    component.timeslot = {} as any;
    component.submitRemove(1);
    expect(component.removeRequest.emit).toHaveBeenCalled();
  });

  it('submitRemove should not emit event if timeslot is undefined', () => {
    spyOn(component.removeRequest, 'emit');
    component.submitRemove(1);
    expect(component.removeRequest.emit).not.toHaveBeenCalled();
  });

  it('submitAdd should emit event', () => {
    spyOn(component.addRequest, 'emit');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    component.timeslot = {} as any;
    component.formGroup.controls['createForm'].setValue({
      weekday: 1,
      startTime: '08:00',
      endTime: '16:00',
    });

    fixture.detectChanges();

    component.submitAdd();
    expect(component.addRequest.emit).toHaveBeenCalled();
  });

  it('submitAdd should not emit event if timeslot is undefined', () => {
    spyOn(component.addRequest, 'emit');
    component.submitAdd();
    expect(component.addRequest.emit).not.toHaveBeenCalled();
  });

  it('submitAdd should not emit event if formGroup is invalid', () => {
    spyOn(component.addRequest, 'emit');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    component.timeslot = {} as any;

    component.formGroup.controls['createForm'].setErrors({ pattern: true });
    component.submitAdd();
    expect(component.addRequest.emit).not.toHaveBeenCalled();
  });

  it('should patch formGroup on ngOnChanges', () => {
    spyOn(component.formGroup, 'patchValue');
    component.ngOnChanges({
      timeslot: {
        currentValue: {
          weekdays: [
            {
              start_time: '08:00',
              end_time: '16:00',
            },
          ],
        },
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    fixture.detectChanges();
    expect(component.formGroup.patchValue).toHaveBeenCalled();
  });

  it('should not patch formGroup on ngOnChanges if timeslot is undefined', () => {
    spyOn(component.formGroup, 'patchValue');
    component.ngOnChanges({
      timeslot: {
        currentValue: undefined,
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    fixture.detectChanges();
    expect(component.formGroup.patchValue).not.toHaveBeenCalled();
  });

  it('should initialize validWeekdayOptions on ngOnChanges with all weekdays', () => {
    component.timeslot = {
      weekdays: [
        {
          id: 1,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 2,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 3,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 4,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 5,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 6,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 7,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
      ],
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any;
    component.ngOnChanges({
      timeslot: {
        currentValue: {
          weekdays: [
            {
              id: 1,
            },
            {
              id: 2,
            },
            {
              id: 3,
            },
            {
              id: 4,
            },
            {
              id: 5,
            },
            {
              id: 6,
            },
            {
              id: 7,
            },
          ],
        },
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    fixture.detectChanges();
    expect(component.validWeekdayOptions.length).toEqual(0);
  });
  it('should initialize validWeekdayOptions on ngOnChanges with more weekdays', () => {
    component.timeslot = {
      weekdays: [
        {
          id: 1,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 2,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 3,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 4,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 5,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
        {
          id: 6,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
      ],
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any;
    component.ngOnChanges({
      timeslot: {
        currentValue: {
          weekdays: [
            {
              id: 1,
            },
            {
              id: 2,
            },
            {
              id: 3,
            },
            {
              id: 4,
            },
            {
              id: 5,
            },
            {
              id: 6,
            },
          ],
        },
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    fixture.detectChanges();
    expect(component.validWeekdayOptions.length).toEqual(1);
    expect(component.validWeekdayOptions[0].id).toEqual(7);
  });
  it('should initialize validWeekdayOptions on ngOnChanges with one', () => {
    component.timeslot = {
      weekdays: [
        {
          id: 1,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
      ],
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any;
    component.ngOnChanges({
      timeslot: {
        currentValue: {
          weekdays: [
            {
              id: 1,
            },
          ],
        },
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    fixture.detectChanges();
    expect(component.validWeekdayOptions.length).toEqual(6);
    expect(component.validWeekdayOptions[0].id).toEqual(2);
    expect(component.validWeekdayOptions[1].id).toEqual(3);
    expect(component.validWeekdayOptions[2].id).toEqual(4);
    expect(component.validWeekdayOptions[3].id).toEqual(5);
    expect(component.validWeekdayOptions[4].id).toEqual(6);
    expect(component.validWeekdayOptions[5].id).toEqual(7);
  });

  it('should initialize weekdays formArray with 7 elements', () => {
    component.timeslot = {
      weekdays: [
        {
          id: 1,
          start_time: '08:00',
          end_time: '16:00',
          name: 'test',
        },
      ],
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any;
    component.ngOnChanges({
      timeslot: {
        currentValue: {
          weekdays: [
            {
              id: 1,
            },
          ],
        },
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } as any);
    fixture.detectChanges();
    expect(component.formGroup.controls['weekdays'].value.length).toEqual(7);
  });

  it('should render Bitte einen Timeslot auswählen if timeslot is null', () => {
    component.timeslot = null;
    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('h3').textContent).toContain('Bitte einen Timeslot auswählen');
  });

  it('should make the control required', () => {
    component.formGroup.controls.createForm.setValue({
      weekday: 1,
      startTime: '08:00',
      endTime: '16:00',
    });

    expect(component.formGroup.controls.createForm.valid).toBeTrue();

    component.formGroup.controls.createForm.setValue({
      weekday: 1,
      startTime: '08:00',
      endTime: '',
    });
    expect(component.formGroup.controls.createForm.valid).toBeFalse();

    component.formGroup.controls.createForm.setValue({
      weekday: 1,
      startTime: '',
      endTime: '16:00',
    });
    expect(component.formGroup.controls.createForm.valid).toBeFalse();

    component.formGroup.controls.createForm.setValue({
      weekday: 1,
      startTime: '',
      endTime: '',
    });
    expect(component.formGroup.controls.createForm.valid).toBeFalse();
  });

  it('add-button should be disabled if formGroup is invalid', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    component.timeslot = {} as any;

    component.formGroup.controls.createForm.setValue({
      weekday: 1,
      startTime: '08:00',
      endTime: '16:00',
    });

    fixture.detectChanges();
    const compiled = fixture.nativeElement;
    expect(compiled.querySelector('#add-button').disabled).toBeFalse();

    component.formGroup.controls.createForm.setValue({
      weekday: 1,
      startTime: '08:00',
      endTime: '',
    });

    fixture.detectChanges();
    expect(compiled.querySelector('#add-button').disabled).toBeTrue();
  });
});
