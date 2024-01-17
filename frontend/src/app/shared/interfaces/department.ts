interface Metadata {
  created_at: Date;
  updated_at: Date;
  deleted_at: Date | null;
}

export interface Department {
  name: string;
}

export type DetailedDepartmentWithMetadata = Department & Metadata;

export type CreateDepartment = Pick<Department, 'name'>;
