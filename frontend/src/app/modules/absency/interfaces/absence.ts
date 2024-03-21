/** type used to keep track of the absences for a person **/
export type AbsenceForPerson = {
  startDate: Date;
  endDate: Date;
  name: string;
  created_at: Date;
  color: string;
};

/** type used to keep track of the absences for a date e.g.: each date in a week **/
export type AbsenceByDate = {
  personID: string;
  reason: string;
  date: string;
  createdAt: Date;
};

export type AbsenceReponse = {
  person_id: string;
  reason: string;
  date: string;
  created_at: Date;
};
