import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';

export type DisplayedWorkdayTimeslot = WorkdayTimeslot & {
  // for display
  gridColumn: number;

  // color for the light and dark mode
  colorForLightMode: string | null;
  colorForDarkMode: string | null;

  // check if start_time and end_time is 00:00, if so then we dont want to display the time, thus we store a boolean to check if it is valid
  validTime: boolean;
};

export type DisplayedWorkdayTimeslotGroup = {
  name: string;
  id: string;
  workdayTimeslots: DisplayedWorkdayTimeslot[];

  // This is a concatenation of the start and end time of the first timeslot in the group.
  startTime: string;
  endTime: string;

  // for display
  gridRow: number;
};

export type DisplayedWorkplace = {
  workplace: WorkplaceWithMetadata;
  name: string;
  timeslotGroups: DisplayedWorkdayTimeslotGroup[];

  // for display
  gridRowStart: number;
  gridRowEnd: number;
};
