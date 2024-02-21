import { TestBed } from '@angular/core/testing';

import { ActiveWeekHandlerService } from './active-week-handler.service';

describe('ActiveWeekHandlerService', () => {
  let service: ActiveWeekHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ActiveWeekHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should have activeWeek$', () => {
    expect(service.activeWeek$).toBeTruthy();
  });

  it('should correctly set the week upon date input 2021-06-07 (MON)', () => {
    const date = new Date('2021-06-07');
    service.activeWeekByDate = date;

    expect(service.activeWeek$).toEqual([
      { name: 'Montag', shortName: 'Mo', date, dateString: '2021-06-07' },
      { name: 'Dienstag', shortName: 'Di', date: new Date('2021-06-08'), dateString: '2021-06-08' },
      { name: 'Mittwoch', shortName: 'Mi', date: new Date('2021-06-09'), dateString: '2021-06-09' },
      { name: 'Donnerstag', shortName: 'Do', date: new Date('2021-06-10'), dateString: '2021-06-10' },
      { name: 'Freitag', shortName: 'Fr', date: new Date('2021-06-11'), dateString: '2021-06-11' },
      { name: 'Samstag', shortName: 'Sa', date: new Date('2021-06-12'), dateString: '2021-06-12' },
      { name: 'Sonntag', shortName: 'So', date: new Date('2021-06-13'), dateString: '2021-06-13' },
    ]);
  });
  it('should correctly set the week upon date input 2021-06-08 (TUE)', () => {
    const date = new Date('2021-06-08');
    service.activeWeekByDate = date;

    expect(service.activeWeek$).toEqual([
      { name: 'Montag', shortName: 'Mo', date: new Date('2021-06-07'), dateString: '2021-06-07' },
      { name: 'Dienstag', shortName: 'Di', date, dateString: '2021-06-08' },
      { name: 'Mittwoch', shortName: 'Mi', date: new Date('2021-06-09'), dateString: '2021-06-09' },
      { name: 'Donnerstag', shortName: 'Do', date: new Date('2021-06-10'), dateString: '2021-06-10' },
      { name: 'Freitag', shortName: 'Fr', date: new Date('2021-06-11'), dateString: '2021-06-11' },
      { name: 'Samstag', shortName: 'Sa', date: new Date('2021-06-12'), dateString: '2021-06-12' },
      { name: 'Sonntag', shortName: 'So', date: new Date('2021-06-13'), dateString: '2021-06-13' },
    ]);
  });
  it('should correctly set the week upon date input  2021-06-13 (SUN)', () => {
    const date = new Date('2021-06-13');
    service.activeWeekByDate = date;

    expect(service.activeWeek$).toEqual([
      { name: 'Montag', shortName: 'Mo', date: new Date('2021-06-07'), dateString: '2021-06-07' },
      { name: 'Dienstag', shortName: 'Di', date: new Date('2021-06-08'), dateString: '2021-06-08' },
      { name: 'Mittwoch', shortName: 'Mi', date: new Date('2021-06-09'), dateString: '2021-06-09' },
      { name: 'Donnerstag', shortName: 'Do', date: new Date('2021-06-10'), dateString: '2021-06-10' },
      { name: 'Freitag', shortName: 'Fr', date: new Date('2021-06-11'), dateString: '2021-06-11' },
      { name: 'Samstag', shortName: 'Sa', date: new Date('2021-06-12'), dateString: '2021-06-12' },
      { name: 'Sonntag', shortName: 'So', date, dateString: '2021-06-13' },
    ]);
  });

  it('should correctly shift the week forward by 1 week', () => {
    // set the date
    const date = new Date('2021-06-07');
    service.activeWeekByDate = date;

    // shift the week by 1 week
    service.shiftWeek(1);

    expect(service.activeWeek$).toEqual([
      { name: 'Montag', shortName: 'Mo', date: new Date('2021-06-14'), dateString: '2021-06-14' },
      { name: 'Dienstag', shortName: 'Di', date: new Date('2021-06-15'), dateString: '2021-06-15' },
      { name: 'Mittwoch', shortName: 'Mi', date: new Date('2021-06-16'), dateString: '2021-06-16' },
      { name: 'Donnerstag', shortName: 'Do', date: new Date('2021-06-17'), dateString: '2021-06-17' },
      { name: 'Freitag', shortName: 'Fr', date: new Date('2021-06-18'), dateString: '2021-06-18' },
      { name: 'Samstag', shortName: 'Sa', date: new Date('2021-06-19'), dateString: '2021-06-19' },
      { name: 'Sonntag', shortName: 'So', date: new Date('2021-06-20'), dateString: '2021-06-20' },
    ]);
  });

  it('should correctly shift the week backwards by 1 week', () => {
    // set the date
    const date = new Date('2021-06-07');
    service.activeWeekByDate = date;

    // shift the week by -1 week
    service.shiftWeek(-1);

    expect(service.activeWeek$).toEqual([
      { name: 'Montag', shortName: 'Mo', date: new Date('2021-05-31'), dateString: '2021-05-31' },
      { name: 'Dienstag', shortName: 'Di', date: new Date('2021-06-01'), dateString: '2021-06-01' },
      { name: 'Mittwoch', shortName: 'Mi', date: new Date('2021-06-02'), dateString: '2021-06-02' },
      { name: 'Donnerstag', shortName: 'Do', date: new Date('2021-06-03'), dateString: '2021-06-03' },
      { name: 'Freitag', shortName: 'Fr', date: new Date('2021-06-04'), dateString: '2021-06-04' },
      { name: 'Samstag', shortName: 'Sa', date: new Date('2021-06-05'), dateString: '2021-06-05' },
      { name: 'Sonntag', shortName: 'So', date: new Date('2021-06-06'), dateString: '2021-06-06' },
    ]);
  });
});
