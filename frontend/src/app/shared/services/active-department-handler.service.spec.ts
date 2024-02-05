import { TestBed } from '@angular/core/testing';

import { ActiveDepartmentHandlerService } from './active-department-handler.service';

describe('ActiveDepartmentHandlerService', () => {
  let service: ActiveDepartmentHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ActiveDepartmentHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
