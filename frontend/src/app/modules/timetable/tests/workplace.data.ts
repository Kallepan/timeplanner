import { Workplace } from '../interfaces/timetable.interface';

export const DUMMY_WORKPLACE_DATAS: Workplace[] = [
  {
    name: 'Anlern',
    slots: [
      {
        name: 'A1',
        timeslots: [
          {
            startTime: '08:00',
            endTime: '12:00',
            occupied: false,
            occupiedBy: null,
            gridColumn: 3,
          },
          {
            startTime: '13:00',
            endTime: '17:00',
            occupied: false,
            occupiedBy: null,
            gridColumn: 4,
          },
        ],
        gridRow: 2,
      },
    ],
    gridRowStart: 2,
    gridRowEnd: 3,
  },
  {
    name: 'Arzt',
    slots: [
      {
        name: 'V1',
        timeslots: [],
        gridRow: 4,
      },
      {
        name: 'V2',
        timeslots: [],
        gridRow: 5,
      },
      {
        name: 'V3',
        timeslots: [],
        gridRow: 6,
      },
    ],
    gridRowStart: 4,
    gridRowEnd: 7,
  },
  {
    name: 'Atemwege',
    slots: [
      {
        name: 'A1',
        timeslots: [],
        gridRow: 8,
      },
      {
        name: 'A2',
        timeslots: [
          {
            startTime: '08:00',
            endTime: '12:00',
            occupied: false,
            occupiedBy: null,
            gridColumn: 6,
          },
          {
            startTime: '13:00',
            endTime: '17:00',
            occupied: false,
            occupiedBy: null,
            gridColumn: 5,
          },
        ],
        gridRow: 9,
      },
    ],
    gridRowStart: 8,
    gridRowEnd: 10,
  },
];
