import { ObjectsWithDatesToBeSorted, groupDatesToRanges } from './group-dates-to-ranges.function';

describe('groupDatesToRanges', () => {
  it('should group dates to ranges. linear', () => {
    const dates: ObjectsWithDatesToBeSorted[] = [
      {
        date: new Date('2022-01-01'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-02'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-03'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-04'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-05'),
        created_at: new Date(),
      },
    ];
    const result = groupDatesToRanges(dates, 'name');
    expect(result).toEqual([
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-05'),
      }),
    ]);
  });

  it('should group dates to ranges. non-linear', () => {
    const dates: ObjectsWithDatesToBeSorted[] = [
      {
        date: new Date('2022-01-01'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-02'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-04'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-05'),
        created_at: new Date(),
      },
    ];
    const result = groupDatesToRanges(dates, 'name');
    expect(result).toEqual([
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-02'),
      }),
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-04'),
        endDate: new Date('2022-01-05'),
      }),
    ]);
  });

  it('should group dates to ranges. non-linear, multiple groups', () => {
    const dates: ObjectsWithDatesToBeSorted[] = [
      {
        date: new Date('2022-01-01'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-02'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-04'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-05'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-07'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-08'),
        created_at: new Date(),
      },
    ];
    const result = groupDatesToRanges(dates, 'name');
    expect(result).toEqual([
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-02'),
      }),
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-04'),
        endDate: new Date('2022-01-05'),
      }),
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-07'),
        endDate: new Date('2022-01-08'),
      }),
    ]);
  });

  it('should group single date to range', () => {
    const dates: ObjectsWithDatesToBeSorted[] = [
      {
        date: new Date('2022-01-01'),
        created_at: new Date(),
      },
    ];
    const result = groupDatesToRanges(dates, 'name');
    expect(result).toEqual([
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-01'),
      }),
    ]);
  });
});
