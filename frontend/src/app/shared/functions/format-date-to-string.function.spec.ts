import { formatDateToDateString } from './format-date-to-string.function';

describe('formatDateToDateString', () => {
  it('should format date to string', () => {
    const dates = [
      { date: new Date('2022-01-01'), result: '2022-01-01' },
      { date: new Date('2022-01-02'), result: '2022-01-02' },
      { date: new Date('2022-01-03'), result: '2022-01-03' },
      { date: new Date('2022-01-04'), result: '2022-01-04' },
      { date: new Date('2022-01-05'), result: '2022-01-05' },
      { date: new Date('2022-01-12'), result: '2022-01-12' },
      { date: new Date('2022-12-15'), result: '2022-12-15' },
      { date: new Date('2022-12-16'), result: '2022-12-16' },
    ];

    dates.forEach((d) => {
      expect(formatDateToDateString(d.date)).toEqual(d.result);
    });
  });
});
