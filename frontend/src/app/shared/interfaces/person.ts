interface Metadata {
  created_at: Date;
  updated_at: Date;
  deleted_at: Date | null;
}

export interface Person {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  active: boolean;
  working_hours: number;
  workplaces: string[];
  departments: string[];
  weekdays: string[];
}

export type SimplePerson = Pick<
  Person,
  'id' | 'first_name' | 'last_name' | 'email' | 'active' | 'working_hours'
>;

export type SimplePersonWithMetadata = SimplePerson & Metadata;

export type DetailedPersonWithMetadata = Person & Metadata;
