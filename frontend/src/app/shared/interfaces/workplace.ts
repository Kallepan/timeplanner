import { Metadata } from './base';

export interface Workplace {
  id: string;
  name: string;
  department_id: string;
}

export type WorkplaceWithMetadata = Workplace & Metadata;

export type CreateWorkplace = Pick<Workplace, 'name' | 'id'>;
