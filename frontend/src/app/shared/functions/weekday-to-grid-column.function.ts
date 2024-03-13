import { dateToWeekdayID } from './date-to-weekday-id.function';

export const DateToGridColumn = (d: Date): number => {
  /**
   * Converts a weekday to a grid column number
   *
   * @param weekday The weekday to convert
   * @returns The grid column number
   * @example
   * weekdayToGridColumn(new Date(2023, 1, 2)) // Monday -> 3
   * weekdayToGridColumn(new Date(2023, 1, 3)) // Tuesday -> 4
   **/

  // Get the weekday ID
  const weekdayID = dateToWeekdayID(d);

  return WeekdayIDToGridColumn(weekdayID);
};

export const WeekdayIDToGridColumn = (weekdayID: number): number => {
  /**
   * Converts a weekday ID to a grid column number
   *
   * @param weekdayID The weekday ID to convert
   * @returns The grid column number
   * @example
   * weekdayToGridColumn(1) // Monday -> 3
   * weekdayToGridColumn(2) // Tuesday -> 4
   **/

  if (weekdayID > 7 || weekdayID < 1) {
    throw new Error('Weekday ID cannot be greater than 7 or less than 1');
  }

  return weekdayID + 2;
};
