import { dateToWeekdayID } from './date-to-weekday-id.function';

describe('dateToWeekdayId', () => {
  it('should return 1 for 2021-01-04', () => {
    const date = new Date('2021-01-04');
    expect(dateToWeekdayID(date)).toBe(1);
  });

  it('should return 2 for 2021-01-05', () => {
    const date = new Date('2021-01-05');
    expect(dateToWeekdayID(date)).toBe(2);
  });

  it('should return 3 for 2021-01-06', () => {
    const date = new Date('2021-01-06');
    expect(dateToWeekdayID(date)).toBe(3);
  });

  it('should return 4 for 2020-12-31', () => {
    const date = new Date('2020-12-31');
    expect(dateToWeekdayID(date)).toBe(4);
  });

  it('should return 5 for 2021-01-01', () => {
    const date = new Date('2021-01-01');
    expect(dateToWeekdayID(date)).toBe(5);
  });

  it('should return 6 for 2021-01-02', () => {
    const date = new Date('2021-01-02');
    expect(dateToWeekdayID(date)).toBe(6);
  });

  it('should return 7 for 2021-01-03', () => {
    const date = new Date('2021-01-03');
    expect(dateToWeekdayID(date)).toBe(7);
  });
});
