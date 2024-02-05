import { Injectable, signal } from '@angular/core';

export type Weekday = {
  name: string;
  shortName: string;
  date: Date;
  dateString: string;
};

@Injectable({
  providedIn: 'root',
})
export class ActiveWeekHandlerService {
  // getter and setter for the active week signal
  private _activeWeek = signal<Weekday[]>([]);
  set activeWeekByDate(date: Date) {
    this._activeWeek.set(this._getWeekFromDate(date));
  }
  get activeWeek$() {
    return this._activeWeek();
  }

  constructor() {
    // get dates from monday to sunday for the current week. This is the default active week
    this._activeWeek.set(this._getWeekFromDate(new Date()));
  }

  private _getWeekFromDate(date: Date) {
    // Create a new date object from the input date
    const currentDate = new Date(date.getTime());

    // Calculate the nearest Monday
    while (currentDate.getDay() !== 1) {
      currentDate.setDate(currentDate.getDate() - 1);
    }

    // Create an array to store all dates of the week
    const week = [new Date(currentDate.getTime())];

    // Add the next 6 days to the array
    for (let i = 1; i <= 6; i++) {
      currentDate.setDate(currentDate.getDate() + 1);
      week.push(new Date(currentDate.getTime()));
    }

    // format the Date[] to Weekday[]
    const formattedWeek = week.map((date) => ({
      name: date.toLocaleString('de-DE', { weekday: 'long' }),
      shortName: date.toLocaleString('de-DE', { weekday: 'short' }),
      date,
      dateString: date.toISOString().split('T')[0],
    }));

    return formattedWeek;
  }

  shiftWeek(shift: number) {
    /** Shifts the actively viewed week by + or - x weeks */
    // get the current active week
    const currentWeek = this._activeWeek();

    // get the first date of the current week
    const firstDate = currentWeek[0];

    // create a new date object from the first date
    const newDate = new Date(firstDate.date.getTime());

    // add the shift to the new date
    newDate.setDate(newDate.getDate() + shift * 7);

    // set the active week to the new date
    this._activeWeek.set(this._getWeekFromDate(newDate));
  }
}
