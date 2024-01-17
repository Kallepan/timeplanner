import { Person } from './person';

export type Workday = {
  department: string;
  workplace: string;
  timeslot: string;
  date: string;

  start_time: string;
  end_time: string;

  person: Person | null;
};

export type AssignPersonToWorkdayRequest = {
  date: string;
  department_name: string;
  workplace_name: string;
  timeslot_name: string;
  person_id: string;
};

export type UnassignPersonFromWorkdayRequest = {
  date: string;
  department_name: string;
  workplace_name: string;
  timeslot_name: string;
  person_id: string;
};
