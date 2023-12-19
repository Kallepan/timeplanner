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
  occupied: boolean;
  // Todo: This is only for debug, in the future, this should be a Person array
  occupiedBy: string;
  disabled?: boolean;

  // display
  gridColumn: number;
}

export interface Person {
  email: string;
  firstname: string;
  lastname: string;
  id: string; // 4 Letter identifier

  workingHours: number;
  actualHours: number;
}
