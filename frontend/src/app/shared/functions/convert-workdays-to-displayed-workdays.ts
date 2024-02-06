import { DisplayedWorkdayTimeslot, DisplayedWorkdayTimeslotGroup, DisplayedWorkplace } from '@app/modules/viewer/interfaces/workplace';
import { WorkdayTimeslot } from '../interfaces/workday_timeslot';
import { startTimeToColorForDarkMode, startTimeToColorForLightMode } from './start-time-to-color';
import { WeekdayIDToGridColumn } from './weekday-to-grid-column.function';

export const convertWorkdaysToDisplayedWorkday = (workdays: WorkdayTimeslot[]) => {
  // This counter is used to keep track of the current grid row to be added to the final elements
  let gridRowCounter = 2;

  // The workdays are first grouped by workplace, as each workplace is displayed in a separate row.
  // This is done using a Map where the key is the workplace id and the value is an array of workday timeslots for that workplace.
  const workdayTimeslotGroupedByWorkplaceMap = new Map<string, WorkdayTimeslot[]>();
  workdays.forEach((workdayTimeslot) => {
    const workplaceTimeslots = workdayTimeslotGroupedByWorkplaceMap.get(workdayTimeslot.workplace.id) || [];
    workplaceTimeslots.push(workdayTimeslot);
    workdayTimeslotGroupedByWorkplaceMap.set(
      workdayTimeslot.workplace.id,
      workplaceTimeslots.sort((a, b) => a.timeslot.name.localeCompare(b.timeslot.name)),
    );
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
        colorForLightMode: startTimeToColorForLightMode(workdayTimeslot.start_time),
        colorForDarkMode: startTimeToColorForDarkMode(workdayTimeslot.start_time),
      });
      displayWorkdayTimeslotMap.set(workdayTimeslot.timeslot.name, timeslotsArray);
    });

    // The timeslot groups are then converted to the `DisplayedWorkdayTimeslotGroup` interface.
    // Each group is assigned a grid row number, which is incremented for each group.
    const timeslotGroups: DisplayedWorkdayTimeslotGroup[] = [];
    displayWorkdayTimeslotMap.forEach((displayWorkdayTimeslots, timeslotName) => {
      timeslotGroups.push({
        name: timeslotName,
        workdayTimeslots: displayWorkdayTimeslots,
        gridRow: gridRowCounter,

        id: displayWorkdayTimeslots[0].timeslot.id,
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
      name: workdayTimeslots[0].workplace.name,
      timeslotGroups: timeslotGroups,
      gridRowStart: minGridRow,
      gridRowEnd: maxGridRow,
    });

    // 'reset' the counter for the next i.e. where it should start
    gridRowCounter = maxGridRow + 1;
  });
  // The workplace groups are finally set as the workplaces to be displayed.

  return workplaceGroups;
};
