import { Metadata } from './base';

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

export type PersonWithMetadata = Person & Metadata;

export type CreatePerson = Pick<
  Person,
  'first_name' | 'last_name' | 'email' | 'active' | 'working_hours'
>;
