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
        color: '#ffffff',
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
        color: '#ffffff',
      }),
      jasmine.objectContaining({
        name: 'name',
        startDate: new Date('2022-01-04'),
        endDate: new Date('2022-01-05'),
        color: '#ffffff',
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
    const result = groupDatesToRanges(dates, 'Urlaub');
    expect(result).toEqual([
      jasmine.objectContaining({
        name: 'Urlaub',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-01'),
        color: '#2C8FC9',
      }),
    ]);
  });

  it('should set the colors according to the reason', () => {
    const dates: ObjectsWithDatesToBeSorted[] = [
      {
        date: new Date('2022-01-01'),
        created_at: new Date(),
      },
      {
        date: new Date('2022-01-02'),
        created_at: new Date(),
      },
    ];
    const result = groupDatesToRanges(dates, 'Fortbildung');
    expect(result).toEqual([
      jasmine.objectContaining({
        name: 'Fortbildung',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-02'),
        color: '#B56CE2',
      }),
    ]);

    const result2 = groupDatesToRanges(dates, 'Krank (AU)');
    expect(result2).toEqual([
      jasmine.objectContaining({
        name: 'Krank (AU)',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-02'),
        color: '#9CB703',
      }),
    ]);

    const result3 = groupDatesToRanges(dates, 'Krank (keine AU)');
    expect(result3).toEqual([
      jasmine.objectContaining({
        name: 'Krank (keine AU)',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-02'),
        color: '#F5BB00',
      }),
    ]);

    const result4 = groupDatesToRanges(dates, 'Urlaub');
    expect(result4).toEqual([
      jasmine.objectContaining({
        name: 'Urlaub',
        startDate: new Date('2022-01-01'),
        endDate: new Date('2022-01-02'),
        color: '#2C8FC9',
      }),
    ]);
  });
});
