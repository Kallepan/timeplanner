export const dateToWeekdayID = (date: Date): string => {
  /**
   * Converts a date string to a weekday ID
   *
   * @param date The date string to convert
   * @returns The weekday ID
   * @example
   * dateToWeekdayID('2021-01-01') // 'FRI'
   * dateToWeekdayID('2021-01-02') // 'SAT'
   * dateToWeekdayID('2021-01-03') // 'SUN'
   */
  return date.toLocaleDateString('en-US', { weekday: 'short' }).toUpperCase();
};
