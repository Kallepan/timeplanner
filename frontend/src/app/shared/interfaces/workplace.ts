import { Metadata } from './base';

export interface Workplace {
  id: string;
  name: string;
}

export type WorkplaceWithMetadata = Workplace & Metadata;

export type CreateWorkplace = Pick<Workplace, 'name' | 'id'>;
