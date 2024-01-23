import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';

export type DisplayedWorkdayTimeslot = WorkdayTimeslot & {
  gridColumn: number;
};

export type DisplayedWorkdayTimeslotGroup = {
  name: string;
  workdayTimeslots: DisplayedWorkdayTimeslot[];

  // for display
  gridRow: number;
};

export type DisplayedWorkplace = {
  workplace: WorkplaceWithMetadata;
  timeslotGroups: DisplayedWorkdayTimeslotGroup[];

  // for display
  gridRowStart: number;
  gridRowEnd: number;
};
