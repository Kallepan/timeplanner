import { DepartmentWithMetadata } from '@app/shared/interfaces/department';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';

export const mockActiveWeek = {
  department: 'department1',
  dates: [new Date(2022, 1, 1), new Date(2022, 1, 2)],
};

export const mockDepartment: DepartmentWithMetadata = {
  id: 'dep',
  name: 'department1',
  created_at: new Date(),
  updated_at: new Date(),
  deleted_at: null,
};

export const mockWorkplace: WorkplaceWithMetadata = {
  id: 'workplace1',
  name: 'workplace1',
  created_at: new Date(),
  updated_at: new Date(),
  deleted_at: null,
};

export const mockTimeslot: TimeslotWithMetadata = {
  id: 'timeslot1',
  name: 'timeslot1',
  active: true,
  department_name: 'department1',
  workplace_name: 'workplace1',
  weekdays: [
    {
      id: 'weekday1',
      name: 'Monday',
      start_time: '08:00:00',
      end_time: '16:00:00',
    },
  ],
  created_at: new Date(),
  updated_at: new Date(),
  deleted_at: null,
};

// create a mock workday
export const mockWorkdays: WorkdayTimeslot[] = [
  {
    date: '2022-02-01',
    department: mockDepartment,
    workplace: mockWorkplace,
    timeslot: mockTimeslot,
    start_time: '08:00:00',
    end_time: '16:00:00',
    person: null,
    weekday: 'MON',
    duration_in_minutes: 60,
  },
  {
    date: '2022-02-01',
    department: mockDepartment,
    workplace: mockWorkplace,
    timeslot: mockTimeslot,
    start_time: '08:00:00',
    end_time: '16:00:00',
    person: null,
    weekday: 'MON',
    duration_in_minutes: 60,
  },
];
