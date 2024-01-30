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

  person: PersonWithMetadata | null;
};

export type AssignPersonToWorkdayTimeslotRequest = {
  date: string;
  department_id: string;
  workplace_id: string;
  timeslot_name: string;
  person_id: string;
};

export type UnassignPersonFromWorkdayTimeslotRequest = {
  date: string;
  department_id: string;
  workplace_id: string;
  timeslot_name: string;
  person_id: string;
};
