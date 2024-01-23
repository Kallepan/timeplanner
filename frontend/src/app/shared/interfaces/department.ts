import { Metadata } from './base';

export interface Department {
  id: string;
  name: string;
}

export type DepartmentWithMetadata = Department & Metadata;

export type CreateDepartment = Pick<Department, 'name' | 'id'>;
