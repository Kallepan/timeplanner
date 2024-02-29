import { DepartmentWithMetadata } from '@app/shared/interfaces/department';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { Weekday } from '@app/shared/services/active-week-handler.service';

export const mockActiveWeek: Weekday[] = [
  {
    name: 'Monday',
    shortName: 'Mo',
    date: new Date('2024-01-01'),
    dateString: '2024-01-01',
  },
  {
    name: 'Tuesday',
    shortName: 'Tu',
    date: new Date('2024-01-02'),
    dateString: '2024-01-02',
  },
  {
    name: 'Wednesday',
    shortName: 'We',
    date: new Date('2024-01-03'),
    dateString: '2024-01-03',
  },
  {
    name: 'Thursday',
    shortName: 'Th',
    date: new Date('2024-01-04'),
    dateString: '2024-01-04',
  },
  {
    name: 'Friday',
    shortName: 'Fr',
    date: new Date('2024-01-05'),
    dateString: '2024-01-05',
  },
  {
    name: 'Saturday',
    shortName: 'Sa',
    date: new Date('2024-01-06'),
    dateString: '2024-01-06',
  },
  {
    name: 'Sunday',
    shortName: 'Su',
    date: new Date('2024-01-07'),
    dateString: '2024-01-07',
  },
];
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
  department_id: 'department1',
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
      id: 1,
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
    persons: [],
    weekday: 1,
    duration_in_minutes: 60,
    comment: '',
  },
  {
    date: '2022-02-01',
    department: mockDepartment,
    workplace: mockWorkplace,
    timeslot: mockTimeslot,
    start_time: '08:00:00',
    end_time: '16:00:00',
    persons: [],
    weekday: 1,
    duration_in_minutes: 60,
    comment: 'comment',
  },
];
