import { Component, EventEmitter, Input, OnChanges, Output, SimpleChanges, inject } from '@angular/core';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';
import { MatExpansionModule } from '@angular/material/expansion';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatSelectModule } from '@angular/material/select';
import { constants } from '@app/core/constants/constants';
import { ValidWeekday } from '@app/core/types/weekday';

@Component({
  selector: 'app-edit-weekdays-of-timeslot',
  standalone: true,
  imports: [MatExpansionModule, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatButtonModule, MatIconModule, MatSelectModule],
  templateUrl: './edit-weekdays-of-timeslot.component.html',
  styleUrl: './edit-weekdays-of-timeslot.component.scss',
})
export class EditWeekdaysOfTimeslotComponent implements OnChanges {
  @Input({ required: true }) timeslot: TimeslotWithMetadata | undefined;

  validWeekdayOptions: ValidWeekday[] = [];
  fb = inject(FormBuilder);
  formGroup = this.fb.group({
    createForm: this.fb.group({
      weekday: this.fb.control(0, [Validators.required]),
      startTime: this.fb.control('', [Validators.required, Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      endTime: this.fb.control('', [Validators.required, Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
    }),
    weekdays: this.fb.array([
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
      this.fb.group({
        startTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
        endTime: this.fb.control('', [Validators.pattern('([01]?[0-9]|2[0-3]):[0-5][0-9]')]),
      }),
    ]),
  });

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['timeslot']) {
      this.timeslot = changes['timeslot'].currentValue;
    }

    if (!this.timeslot || !this.timeslot.weekdays) return;

    this.formGroup.patchValue({
      weekdays: this.timeslot.weekdays.map((weekday) => ({
        startTime: weekday.start_time,
        endTime: weekday.end_time,
      })),
    });

    // set it to all weekdays which are not already in the timeslot
    this.validWeekdayOptions = constants.POSSIBLE_WEEKDAYS.filter((weekday) => !(this.timeslot?.weekdays ?? []).some((tsWeekday) => tsWeekday.id === weekday.id));

    // disable the form if there are no more weekdays to add
    if (this.validWeekdayOptions.length === 0) {
      this.formGroup.controls.createForm.disable();
    } else {
      // dont forget to enable it again
      this.formGroup.controls.createForm.enable();
    }
  }

  @Output() editRequest = new EventEmitter<{ startTimeControl: FormControl; endTimeControl: FormControl; timeslot: TimeslotWithMetadata; weekdayID: number }>();
  submitEdit(startTimeControl: FormControl, endTimeControl: FormControl, weekdayID: number) {
    if (startTimeControl.invalid || !this.timeslot || endTimeControl.invalid) return;
    this.editRequest.emit({ startTimeControl, endTimeControl, timeslot: this.timeslot, weekdayID });
    this.formGroup.controls.weekdays.controls[weekdayID - 1].reset();
  }

  @Output() removeRequest = new EventEmitter<{ id: number; timeslot: TimeslotWithMetadata }>();
  submitRemove(id: number) {
    if (!this.timeslot) return;

    this.removeRequest.emit({ id, timeslot: this.timeslot });
    this.validWeekdayOptions.push({ id, name: constants.POSSIBLE_WEEKDAYS.find((weekday) => weekday.id === id)?.name ?? '' });
  }

  @Output() addRequest = new EventEmitter<{ control: FormGroup; type: string; timeslot: TimeslotWithMetadata }>();
  submitAdd() {
    if (this.formGroup.controls.createForm.invalid || !this.timeslot) return;
    this.addRequest.emit({ control: this.formGroup.controls.createForm, type: 'create', timeslot: this.timeslot });

    this.validWeekdayOptions = this.validWeekdayOptions.filter((weekday) => weekday.id !== this.formGroup.controls.createForm.controls.weekday.value);
    this.formGroup.controls.createForm.reset();
  }
}
