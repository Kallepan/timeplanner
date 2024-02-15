/** Returns a string from a formatted date objects in the form of "YYYY-MM-DD"
 *  @param date - the date object to format
 * @returns the formatted date string
 */
export const formatDateToDateString = (date: Date): string => {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();

  return `${year}-${month < 10 ? `0${month}` : month}-${day < 10 ? `0${day}` : day}`;
};
