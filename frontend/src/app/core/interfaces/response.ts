// Standard HTTP response interface
export interface APIResponse<T> {
  data: T;
  message: string;
  status: number;
}
