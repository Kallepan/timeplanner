import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';

export type DisplayedWorkdayTimeslot = WorkdayTimeslot & {
  gridColumn: number;
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
