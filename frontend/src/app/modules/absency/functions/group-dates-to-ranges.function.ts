import { AbsenceForPerson } from '../interfaces/absence';

export type ObjectsWithDatesToBeSorted = { date: Date; created_at: Date };

export const groupDatesToRanges = (objectsSortedByDate: ObjectsWithDatesToBeSorted[], name: string): AbsenceForPerson[] => {
  const ranges: AbsenceForPerson[] = [];

  let objectWithStartDate = objectsSortedByDate[0];
  let objectWithLastDate = objectsSortedByDate[0];

  for (let i = 1; i < objectsSortedByDate.length; i++) {
    if (objectsSortedByDate[i].date.getTime() - objectWithLastDate.date.getTime() === 24 * 60 * 60 * 1000) {
      objectWithLastDate = objectsSortedByDate[i];
      continue;
    }

    ranges.push({
      name: name,
      startDate: objectWithStartDate.date,
      endDate: objectWithLastDate.date,
      created_at: objectWithStartDate.created_at,
    });
    objectWithStartDate = objectsSortedByDate[i];
    objectWithLastDate = objectsSortedByDate[i];
  }

  ranges.push({
    name: name,
    startDate: objectWithStartDate.date,
    endDate: objectWithLastDate.date,
    created_at: objectWithStartDate.created_at,
  });
  return ranges;
};
