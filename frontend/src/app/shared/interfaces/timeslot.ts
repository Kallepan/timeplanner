import { Metadata } from './base';

interface OfferedOnWeekdays {
  id: number; // 1, 2, 3, 4, 5, 6, 7
  name: string;
  start_time: string;
  end_time: string;
}

export interface Timeslot {
  id: string;
  name: string;
  active: boolean;
  department_name: string;
  workplace_name: string;

  weekdays: OfferedOnWeekdays[] | undefined;
}

export type TimeslotWithMetadata = Timeslot & Metadata;
export type CreateTimeslot = Pick<Timeslot, 'name' | 'id' | 'active'>;
