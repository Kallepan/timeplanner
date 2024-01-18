import { Metadata } from './base';

export interface Workplace {
  name: string;
}

export type WorkplaceWithMetadata = Workplace & Metadata;

export type CreateWorkplace = Pick<Workplace, 'name'>;
