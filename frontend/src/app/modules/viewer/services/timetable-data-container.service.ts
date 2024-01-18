/**
 * This service is respsonsible for providing the data for the timetable.
 *
 * The data is provided from the backend as an array of Workday objects.
 * Each workday represents a single timeslot on a given day to which a person can be assigned to.
 *
 * These are grouped by workplace and then by the respective timeslots to generate the timetable.
 * Here we also populate the gridRowStart and gridRowEnd as well as the gridColumn (timeslot)
 * and gridRow (slotsGroup).
 *
 * The final data objects looks like this:
 * Workplace -> Has many TimeslotGroups -> of which each has many Timeslots.
 * In the html template we simple loop over the workplaces and then over the timeslotGroups and timeslots.
 **/
import { Injectable, signal } from '@angular/core';
import { Weekday } from '../interfaces/weekday';
import {
  DisplayedWorkdayTimeslot,
  DisplayedWorkdayTimeslotGroup,
  DisplayedWorkplace,
} from '../interfaces/workplace';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WeekdayIDToGridColumn } from '@app/shared/functions/weekday-to-grid-column.function';

@Injectable({
  providedIn: null,
})
export class TimetableDataContainerService {
  protected _weekdays = signal<Weekday[]>([]);
  get weekdays$(): Weekday[] {
    return this._weekdays();
  }
  set weekdays(value: Date[]) {
    this._weekdays.set(
      value.map((date) => {
        return {
          name: date.toLocaleString('default', { weekday: 'long' }),
          shortName: date.toLocaleString('default', { weekday: 'short' }),
          date,
        };
      }),
    );
  }

  protected _fullHeight = signal<number>(0);
  get fullHeight$(): number {
    return this._fullHeight();
  }
  set fullHeight(value: number) {
    this._fullHeight.set(value);
  }

  protected _displayTime = signal<boolean>(true);
  get displayTime$(): boolean {
    return this._displayTime();
  }
  set displayTime(value: boolean) {
    this._displayTime.set(value);
  }

  // the data for the timetable
  protected _workplaces = signal<DisplayedWorkplace[]>([]);
  get workplaces$(): DisplayedWorkplace[] {
    return this._workplaces();
  }
  set workdays(workdays: WorkdayTimeslot[]) {
    let counter = 2;

    // group by workplace
    const workplaceMap = new Map<string, WorkdayTimeslot[]>();
    workdays.forEach((workday) => {
      if (!workplaceMap.has(workday.workplace)) {
        workplaceMap.set(workday.workplace, []);
      }
      workplaceMap.get(workday.workplace)!.push(workday);
    });

    // group each workplace by timeslot into timeslotGroups
    const workplaceGroups: DisplayedWorkplace[] = [];
    workplaceMap.forEach((workplace, workplaceName) => {
      const timeslotMap = new Map<string, WorkdayTimeslot[]>();
      workplace.forEach((workday) => {
        if (!timeslotMap.has(workday.timeslot)) {
          timeslotMap.set(workday.timeslot, []);
        }
        timeslotMap.get(workday.timeslot)!.push(workday);
      });

      // convert to DisplayedWorkdayTimeslotGroup interface
      const timeslotGroups: DisplayedWorkdayTimeslotGroup[] = [];
      timeslotMap.forEach((workdays, timeslotName) => {
        const timeslots: DisplayedWorkdayTimeslot[] = [];
        workdays.forEach((workday) => {
          timeslots.push({
            ...workday,
            gridColumn: WeekdayIDToGridColumn(workday.weekday),
          });
          counter++;
        });
        timeslotGroups.push({
          name: timeslotName,
          timeslots,
          gridRow: counter,
        });
      });

      // fetch the min and max gridRow to set the gridRowStart and gridRowEnd
      // of the workplace which can span multiple rows
      const minGridRow = Math.min(
        ...timeslotGroups.map((timeslotGroup) => timeslotGroup.gridRow),
      );
      const maxGridRow =
        Math.max(
          ...timeslotGroups.map((timeslotGroup) => timeslotGroup.gridRow),
        ) + 1;

      workplaceGroups.push({
        name: workplaceName,
        timeslotGroups,
        gridRowStart: minGridRow,
        gridRowEnd: maxGridRow,
      });

      // 'reset' the counter for the next i.e. where it should start
      counter = maxGridRow + 1;
    });

    const fullHeight = Math.max(
      ...workplaceGroups.map((workplace) => workplace.gridRowEnd),
    );
    this._fullHeight.set(fullHeight);

    this._workplaces.set(workplaceGroups);
  }
}
