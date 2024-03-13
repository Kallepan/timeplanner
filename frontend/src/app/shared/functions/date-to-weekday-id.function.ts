export const dateToWeekdayID = (date: Date): number => {
  /**
   * Converts a date string to a weekday ID
   *
   * @param date The date string to convert
   * @returns The weekday ID
   * @example
   * dateToWeekdayID('2021-01-01') // 5
   * dateToWeekdayID('2021-01-02') // 6
   * dateToWeekdayID('2021-01-03') // 7
   */
  const weekday = date.getDay();
  return weekday === 0 ? 7 : weekday;
};
