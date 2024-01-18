import { Metadata } from './base';

interface OfferedOnWeekdays {
  id: string;
  name: string;
  start_time: string;
  end_time: string;
}

export interface Timeslot {
  name: string;
  active: boolean;
  department_name: string;
  workplace_name: string;

  weekdays: OfferedOnWeekdays[];
}

export type TimeslotWithMetadata = Timeslot & Metadata;
export type CreateTimeslot = Pick<Timeslot, 'name' | 'active'>;