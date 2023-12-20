import { HttpClient } from '@angular/common/http';
import { Injectable, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { WeekdayIDToGridColumn } from '@app/shared/functions/weekday-to-grid-column.function';
import { Subject, catchError, map, of, tap } from 'rxjs';
import {
  Slot,
  Timeslot,
  TimeslotResponse,
  Weekday,
  Workplace,
} from '../interfaces/timetable.interface';
import { TimetableAPIService } from './timetable-api.service';

@Injectable({
  providedIn: null,
})
export abstract class AbstractTimetableDataService {
  /**
   * This service is respsonsible for providing the data for the timetable.
   * It is abstract, because the data can be provided from different sources.
   * Therefore, a specific implementation is needed.
   *
   * The data is provided as an array of Timeslots from a source, which is
   * converted to the Workplace interface. Here we also populate the
   * gridRowStart and gridRowEnd as well as the gridColumn (timeslot)
   * and gridRow (slot).
   **/
  private timetableAPIService = inject(TimetableAPIService);

  protected readonly http = inject(HttpClient);

  // fullHeight of the timetable
  protected _fullHeight = signal<number>(0);
  get fullHeight$(): number {
    return this._fullHeight();
  }

  // Wether the timetable should display the time or not
  protected _displayTime = signal<boolean>(false);
  get displayTime$(): boolean {
    return this._displayTime();
  }
  set displayTime$(value: boolean) {
    this._displayTime.set(value);
  }

  // keep track of the weekdays
  protected _weekdays = signal<Weekday[]>([]);
  set weekdays$(startDate: Date) {
    const weekdays: Weekday[] = [];
    for (let i = 0; i < 7; i++) {
      const date = new Date(startDate);
      date.setDate(startDate.getDate() + i);
      weekdays.push({
        name: date.toLocaleString('default', { weekday: 'long' }),
        shortName: date.toLocaleString('default', { weekday: 'short' }),
        date,
      });
    }
    this._weekdays.set(weekdays);
  }
  get weekdays$(): Weekday[] {
    return this._weekdays();
  }

  public _slots = new Subject<TimeslotResponse[]>();
  slots$ = this._slots.asObservable().pipe(
    takeUntilDestroyed(),
    map((timeslots) => {
      let counter = 2;
      // group By workplace
      const workplaceMap = new Map<string, TimeslotResponse[]>();
      timeslots.forEach((timeslot) => {
        if (!workplaceMap.has(timeslot.workplace_name)) {
          workplaceMap.set(timeslot.workplace_name, []);
        }
        workplaceMap.get(timeslot.workplace_name)!.push(timeslot);
      });

      // group by slot
      const workplaceDatas: Workplace[] = [];
      workplaceMap.forEach((timeslots, workplaceName) => {
        const slotMap = new Map<string, TimeslotResponse[]>();
        timeslots.forEach((timeslot) => {
          if (!slotMap.has(timeslot.name)) {
            slotMap.set(timeslot.name, []);
          }
          slotMap.get(timeslot.name)!.push(timeslot);
        });

        // convert to Workplace interface
        const workplaceSlots: Slot[] = [];
        slotMap.forEach((timeslots, slotName) => {
          const slots: Timeslot[] = [];
          timeslots.forEach((timeslot) => {
            slots.push({
              startTime: timeslot.start_time,
              endTime: timeslot.end_time,
              occupied: false,
              occupiedBy: null,
              disabled: timeslot.disabled,
              gridColumn: WeekdayIDToGridColumn(timeslot.weekday_id),
            });
            // increment it by one as each timeslot is displayed in its own row
            counter++;
          });
          workplaceSlots.push({
            name: slotName,
            timeslots: slots,
            gridRow: counter,
          });
        });
        // fetch the min and max gridRow to set the gridRowStart and gridRowEnd
        // of the workplace which can span multiple rows
        const minGridRow = Math.min(
          ...workplaceSlots.map((slot) => slot.gridRow),
        );
        const maxGridRow =
          Math.max(...workplaceSlots.map((slot) => slot.gridRow)) + 1;
        workplaceDatas.push({
          name: workplaceName,
          slots: workplaceSlots,
          gridRowStart: minGridRow,
          gridRowEnd: maxGridRow,
        });

        // 'reset' the counter for the next i.e. where it should start
        counter = maxGridRow + 1;
      });
      return workplaceDatas;
    }),
    tap((workplaceDatas) => {
      // Extract the largest gridRowEnd to set fullHeight of the timetable
      const fullHeight = Math.max(
        ...workplaceDatas.map((workplace) => workplace.gridRowEnd),
      );

      this._fullHeight.set(fullHeight);
    }),
    catchError(() => of([])),
  );
  private _workplaces = signal<Workplace[]>([]);
  get workplaces$(): Workplace[] {
    return this._workplaces();
  }

  constructor() {
    this.slots$.subscribe((workplaces) => {
      this._workplaces.set(workplaces);
    });
  }
}
