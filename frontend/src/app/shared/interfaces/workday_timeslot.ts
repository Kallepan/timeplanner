import { DepartmentWithMetadata } from './department';
import { PersonWithMetadata } from './person';
import { TimeslotWithMetadata } from './timeslot';
import { WorkplaceWithMetadata } from './workplace';

export type WorkdayTimeslot = {
  department: DepartmentWithMetadata;
  workplace: WorkplaceWithMetadata;
  timeslot: TimeslotWithMetadata;
  date: string;

  weekday: string; // MON, TUE, WED, THU, FRI, SAT, SUN

  start_time: string;
  end_time: string;
  duration_in_minutes: number;
  comment: string;

  persons: PersonWithMetadata[];
};

export type AssignPersonToWorkdayTimeslotRequest = {
  date: string;
  department_id: string;
  workplace_id: string;
  timeslot_id: string;
  person_id: string;
};

export type UnassignPersonFromWorkdayTimeslotRequest = {
  date: string;
  department_id: string;
  workplace_id: string;
  timeslot_id: string;
  person_id: string;
};

export type UpdateWorkdayRequest = {
  date: string;
  department_id: string;
  workplace_id: string;
  timeslot_id: string;

  // attributes to update
  start_time: string;
  end_time: string;
  comment: string;
  active: boolean;
};
