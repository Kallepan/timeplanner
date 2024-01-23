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
import { DisplayedWorkdayTimeslot, DisplayedWorkdayTimeslotGroup, DisplayedWorkplace } from '../interfaces/workplace';
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
          name: date.toLocaleString('de-DE', { weekday: 'long' }),
          shortName: date.toLocaleString('de-DE', { weekday: 'short' }),
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

  protected _displayPersons = signal<boolean>(true);
  get displayPersons$(): boolean {
    return this._displayPersons();
  }
  set displayPersons(value: boolean) {
    this._displayPersons.set(value);
  }

  protected _displayTime = signal<boolean>(true);
  get displayTime$(): boolean {
    return this._displayTime();
  }
  set displayTime(value: boolean) {
    this._displayTime.set(value);
  }

  protected _colorize = signal<boolean>(true);
  get colorize$(): boolean {
    return this._colorize();
  }
  set colorize(value: boolean) {
    this._colorize.set(value);
  }

  // the data for the timetable
  protected _workplaces = signal<DisplayedWorkplace[]>([]);
  get workplaces$(): DisplayedWorkplace[] {
    return this._workplaces();
  }
  /**
   * This method is a setter for the `workdays` property. It takes an array of `WorkdayTimeslot` objects as input.
   * Each `WorkdayTimeslot` object represents a timeslot in a workday at a specific workplace.
   * The method organizes these timeslots into a structure that is convenient for displaying in a grid.
   *
   * @param workdays - An array of `WorkdayTimeslot` objects.
   **/
  set workdays(workdays: WorkdayTimeslot[]) {
    // This counter is used to keep track of the current grid row to be added to the final elements
    let gridRowCounter = 2;

    // The workdays are first grouped by workplace, as each workplace is displayed in a separate row.
    // This is done using a Map where the key is the workplace id and the value is an array of workday timeslots for that workplace.
    const workdayTimeslotGroupedByWorkplaceMap = new Map<string, WorkdayTimeslot[]>();
    workdays.forEach((workdayTimeslot) => {
      const workplaceTimeslots = workdayTimeslotGroupedByWorkplaceMap.get(workdayTimeslot.workplace.id) || [];
      workplaceTimeslots.push(workdayTimeslot);
      workdayTimeslotGroupedByWorkplaceMap.set(workdayTimeslot.workplace.id, workplaceTimeslots);
    });

    // The workdays for each workplace are then further grouped by timeslot, as each timeslot is displayed in a separate subrow within the workplace row.
    // This is done using a Map where the key is the timeslot name and the value is an array of workday timeslots for that timeslot.
    const workplaceGroups: DisplayedWorkplace[] = [];
    workdayTimeslotGroupedByWorkplaceMap.forEach((workdayTimeslots) => {
      const displayWorkdayTimeslotMap = new Map<string, DisplayedWorkdayTimeslot[]>();
      workdayTimeslots.forEach((workdayTimeslot) => {
        const timeslotsArray = displayWorkdayTimeslotMap.get(workdayTimeslot.timeslot.name) || [];

        // add the gridColumn property to the workday
        timeslotsArray.push({
          ...workdayTimeslot,
          gridColumn: WeekdayIDToGridColumn(workdayTimeslot.weekday),
        });
        displayWorkdayTimeslotMap.set(workdayTimeslot.timeslot.name, timeslotsArray);
      });

      // The timeslot groups are then converted to the `DisplayedWorkdayTimeslotGroup` interface.
      // Each group is assigned a grid row number, which is incremented for each group.
      const timeslotGroups: DisplayedWorkdayTimeslotGroup[] = [];
      displayWorkdayTimeslotMap.forEach((displayWorkdayTimeslots, timeslotName) => {
        timeslotGroups.push({
          name: timeslotName,
          workdayTimeslots: displayWorkdayTimeslots.sort((a, b) => a.timeslot.name.localeCompare(b.timeslot.name)),
          gridRow: gridRowCounter,

          startTime: displayWorkdayTimeslots[0].start_time,
          endTime: displayWorkdayTimeslots[0].end_time,
        });
        gridRowCounter++;
      });

      // The minimum and maximum grid row numbers are calculated for each workplace.
      // These are used to set the `gridRowStart` and `gridRowEnd` properties of the workplace, which can span multiple rows.
      const minGridRow = Math.min(...timeslotGroups.map((timeslotGroup) => timeslotGroup.gridRow));
      const maxGridRow = Math.max(...timeslotGroups.map((timeslotGroup) => timeslotGroup.gridRow)) + 1; // +1 due to display grid way of handling gridRowEnd

      // The workplace groups are then created, each with its timeslot groups and grid row start and end numbers.
      workplaceGroups.push({
        workplace: workdayTimeslots[0].workplace,
        timeslotGroups: timeslotGroups,
        gridRowStart: minGridRow,
        gridRowEnd: maxGridRow,
      });

      // 'reset' the counter for the next i.e. where it should start
      gridRowCounter = maxGridRow + 1;
    });

    // The maximum grid row end number is calculated and set as the full height of the grid.
    const fullHeight = Math.max(...workplaceGroups.map((workplace) => workplace.gridRowEnd));
    this._fullHeight.set(fullHeight);

    // The workplace groups are finally set as the workplaces to be displayed.
    this._workplaces.set(workplaceGroups);
  }
}
