export interface TimeslotResponse {
  // the timeslot returned from the backend

  department_name: string;
  workplace_name: string;
  name: string;

  weekday_id: string; // 3 Letter identifier
  start_time: string;
  end_time: string;
  disabled?: boolean;

  persons: PersonResponse[];
}

export interface PersonResponse {
  // the person returned from the backend

  email: string;
  firstname: string;
  lastname: string;
  id: string; // 4 Letter identifier

  working_hours: number;
  actual_hours: number;
}
