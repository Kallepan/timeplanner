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

export const WeekdayIDToGridColumn = (weekdayID: string): number => {
  /**
   * Converts a weekday ID to a grid column number
   *
   * @param weekdayID The weekday ID to convert
   * @returns The grid column number
   * @example
   * weekdayToGridColumn('MON') // Monday -> 3
   * weekdayToGridColumn('TUE') // Tuesday -> 4
   **/

  switch (weekdayID) {
    case 'MON':
      return 3;
    case 'TUE':
      return 4;
    case 'WED':
      return 5;
    case 'THU':
      return 6;
    case 'FRI':
      return 7;
    case 'SAT':
      return 8;
    case 'SUN':
      return 9;
    default:
      throw new Error(`Invalid weekday ID: ${weekdayID}`);
  }
};
