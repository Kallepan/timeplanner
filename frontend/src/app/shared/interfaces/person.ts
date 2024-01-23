import { Metadata } from './base';
import { Department } from './department';
import { Workplace } from './workplace';

type OnWeekday = {
  id: string;
  name: string;
};

export interface Person {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  active: boolean;
  working_hours: number;
  workplaces: Workplace[];
  departments: Department[];
  weekdays: OnWeekday[];
}

export type PersonWithMetadata = Person & Metadata;

export type CreatePerson = Pick<Person, 'first_name' | 'last_name' | 'email' | 'active' | 'working_hours'>;
