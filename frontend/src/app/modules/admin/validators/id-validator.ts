import { AsyncValidatorFn } from '@angular/forms';
import { Observable, map } from 'rxjs';

export interface CheckIDExistsInterface {
  checkIDExists(id: string, departmentID?: string, workplaceID?: string): Observable<boolean>;
}

export function AsyncIDValidator<T extends CheckIDExistsInterface>(service: T, departmentID?: string, workplaceID?: string): AsyncValidatorFn {
  return (control) => {
    return service.checkIDExists(control.value, departmentID, workplaceID).pipe(map((res) => (res ? { idExists: true } : null)));
  };
}
