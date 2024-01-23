import { Person, PersonResponse } from './person.interface';

// Stores display information for the weekday
export interface Weekday {
  name: string;
  shortName: string;
  date: Date;
}

export interface Workplace {
  name: string;
  slots: Slot[];

  // display
  gridRowStart: number;
  gridRowEnd: number;
}

export interface Slot {
  name: string;
  timeslots: Timeslot[];

  // display
  gridRow: number;
}

export interface Timeslot {
  startTime: string;
  endTime: string;

  // additional information
  // Todo: This is only for debug, in the future, this should be a Person array
  occupiedBy: Person | null;
  disabled?: boolean;

  // display
  gridColumn: number;
}

// the timeslot returned from the backend
export interface TimeslotResponse {
  department_id: string;
  workplace_id: string;
  name: string;

  weekday_id: string; // 3 Letter identifier
  start_time: string;
  end_time: string;
  disabled?: boolean;

  persons: PersonResponse | null;
}
