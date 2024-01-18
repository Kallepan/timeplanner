import { Metadata } from './base';

export interface Department {
  name: string;
}

export type DepartmentWithMetadata = Department & Metadata;

export type CreateDepartment = Pick<Department, 'name'>;
