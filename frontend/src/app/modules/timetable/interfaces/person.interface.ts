export interface Person {
  email: string;
  firstname: string;
  lastname: string;
  id: string; // 4 Letter identifier

  workingHours: number;
  actualHours: number;

  // calculated by the frontend
  fullname: string;
}

// the person returned from the backend
export interface PersonResponse {
  email: string;
  firstname: string;
  lastname: string;
  id: string; // 4 Letter identifier

  working_hours: number;
  actual_hours: number;
}
