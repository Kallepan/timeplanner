import { constants } from '@app/core/constants/constants';
import { AbsenceForPerson } from '../interfaces/absence';

export type ObjectsWithDatesToBeSorted = { date: Date; created_at: Date };

export const groupDatesToRanges = (objectsSortedByDate: ObjectsWithDatesToBeSorted[], reason: string): AbsenceForPerson[] => {
  const ranges: AbsenceForPerson[] = [];

  let objectWithStartDate = objectsSortedByDate[0];
  let objectWithLastDate = objectsSortedByDate[0];

  for (let i = 1; i < objectsSortedByDate.length; i++) {
    if (objectsSortedByDate[i].date.getTime() - objectWithLastDate.date.getTime() === 24 * 60 * 60 * 1000) {
      objectWithLastDate = objectsSortedByDate[i];
      continue;
    }

    ranges.push({
      name: reason,
      startDate: objectWithStartDate.date,
      endDate: objectWithLastDate.date,
      created_at: objectWithStartDate.created_at,
      color: constants.ABSENCY_REASONS.get(reason) ?? '#ffffff',
    });
    objectWithStartDate = objectsSortedByDate[i];
    objectWithLastDate = objectsSortedByDate[i];
  }

  ranges.push({
    name: reason,
    startDate: objectWithStartDate.date,
    endDate: objectWithLastDate.date,
    created_at: objectWithStartDate.created_at,
    color: constants.ABSENCY_REASONS.get(reason) ?? '#ffffff',
  });
  return ranges;
};
