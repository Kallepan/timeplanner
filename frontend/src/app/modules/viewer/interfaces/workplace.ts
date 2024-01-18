import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';

export type DisplayedWorkdayTimeslot = Omit<
  WorkdayTimeslot,
  'date' | 'department' | 'workplace' | 'timeslot'
> & {
  gridColumn: number;
};

export type DisplayedWorkdayTimeslotGroup = {
  name: string;
  timeslots: DisplayedWorkdayTimeslot[];

  // for display
  gridRow: number;
};

export type DisplayedWorkplace = {
  name: string;
  timeslotGroups: DisplayedWorkdayTimeslotGroup[];

  // for display
  gridRowStart: number;
  gridRowEnd: number;
};
