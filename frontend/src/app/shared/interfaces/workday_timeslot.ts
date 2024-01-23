import { DepartmentWithMetadata } from './department';
import { Person } from './person';
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

  person: WorkdayTimeslotPerson | null;
};

export type AssignPersonToWorkdayTimeslotRequest = {
  date: string;
  department_name: string;
  workplace_name: string;
  timeslot_name: string;
  person_id: string;
};

export type UnassignPersonFromWorkdayTimeslotRequest = {
  date: string;
  department_name: string;
  workplace_name: string;
  timeslot_name: string;
  person_id: string;
};

export type WorkdayTimeslotPerson = Pick<Person, 'id' | 'first_name' | 'last_name' | 'email' | 'working_hours'>;
