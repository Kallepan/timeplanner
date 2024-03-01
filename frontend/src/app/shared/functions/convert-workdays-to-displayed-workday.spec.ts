import { WorkdayTimeslot } from '../interfaces/workday_timeslot';
import { _groupWorkdaysByWorkplace, convertWorkdaysToDisplayedWorkday } from './convert-workdays-to-displayed-workdays';

describe('convertWorkdaysToDisplayedWorkday', () => {
  it('should group workdays by workplace id and timeslot name', () => {
    const workdays: WorkdayTimeslot[] = [
      {
        department: {
          id: '1',
          name: 'Department 1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        workplace: {
          id: '1',
          name: 'Workplace 1',
          department_id: '1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        timeslot: {
          id: '1',
          name: 'Morning',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          active: true,
          department_id: 'Department 1',
          workplace_id: 'Workplace 1',
          weekdays: [],
        },
        date: '2022-01-01',
        weekday: 6,
        start_time: '08:00',
        end_time: '12:00',
        duration_in_minutes: 240,
        comment: '',
        persons: [
          /* PersonWithMetadata[] */
        ],
      },
      {
        department: {
          id: '1',
          name: 'Department 1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        workplace: {
          id: '1',
          name: 'Workplace 1',
          department_id: '1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        timeslot: {
          id: '2',
          name: 'Evening',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          active: true,
          department_id: 'Department 1',
          workplace_id: 'Workplace 1',
          weekdays: [],
        },
        date: '2022-01-01',
        weekday: 6,
        start_time: '13:00',
        end_time: '17:00',
        duration_in_minutes: 240,
        comment: '',
        persons: [
          /* PersonWithMetadata[] */
        ],
      },
    ];

    const displayedWorkdays = convertWorkdaysToDisplayedWorkday(workdays);

    expect(displayedWorkdays.length).toBe(1);
    expect(displayedWorkdays[0].workplace.id).toBe('1');
    expect(displayedWorkdays[0].workplace.name).toBe('Workplace 1');

    expect(displayedWorkdays[0].timeslotGroups.length).toBe(2);
    // different sorting due to alphabetical order
    expect(displayedWorkdays[0].timeslotGroups[0].name).toBe('Evening');
    expect(displayedWorkdays[0].timeslotGroups[1].name).toBe('Morning');

    expect(displayedWorkdays[0].timeslotGroups[0].workdayTimeslots.length).toBe(1);
    expect(displayedWorkdays[0].timeslotGroups[1].workdayTimeslots.length).toBe(1);

    expect(displayedWorkdays[0].timeslotGroups[0].workdayTimeslots[0].timeslot.name).toBe('Evening');
    expect(displayedWorkdays[0].timeslotGroups[1].workdayTimeslots[0].timeslot.name).toBe('Morning');
  });
});

describe('_groupWorkdaysByWorkplace', () => {
  it('should group workdays by workplace id', () => {
    const workdays: WorkdayTimeslot[] = [
      {
        department: {
          id: '1',
          name: 'Department 1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        workplace: {
          id: '1',
          name: 'Workplace 1',
          department_id: '1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        timeslot: {
          id: '1',
          name: 'Morning',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          active: true,
          department_id: 'Department 1',
          workplace_id: 'Workplace 1',
          weekdays: [],
        },
        date: '2022-01-01',
        weekday: 6,
        start_time: '08:00',
        end_time: '12:00',
        duration_in_minutes: 240,
        comment: '',
        persons: [
          /* PersonWithMetadata[] */
        ],
      },
      {
        department: {
          id: '1',
          name: 'Department 1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        workplace: {
          id: '1',
          name: 'Workplace 1',
          department_id: '1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        timeslot: {
          id: '2',
          name: 'Evening',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          active: true,
          department_id: 'Department 1',
          workplace_id: 'Workplace 1',
          weekdays: [],
        },
        date: '2022-01-01',
        weekday: 6,
        start_time: '13:00',
        end_time: '17:00',
        duration_in_minutes: 240,
        comment: '',
        persons: [
          /* PersonWithMetadata[] */
        ],
      },
      {
        department: {
          id: '1',
          name: 'Department 1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        workplace: {
          id: '2',
          name: 'Workplace 2',
          department_id: '1',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
        },
        timeslot: {
          id: '3',
          name: 'Morning',
          created_at: new Date(),
          updated_at: new Date(),
          deleted_at: null,
          active: true,
          department_id: 'Department 1',
          workplace_id: 'Workplace 2',
          weekdays: [],
        },
        date: '2022-01-01',
        weekday: 6,
        start_time: '18:00',
        end_time: '22:00',
        duration_in_minutes: 240,
        comment: '',
        persons: [
          /* PersonWithMetadata[] */
        ],
      },
    ];

    const groupedWorkdays = _groupWorkdaysByWorkplace(workdays);

    expect(groupedWorkdays.get('1')?.length).toBe(2);
    expect(groupedWorkdays.get('2')?.length).toBe(1);
  });
});
