import { TestBed } from '@angular/core/testing';

import { PlannerStateHandlerService } from './planner-state-handler.service';

describe('PlannerStateHandlerService', () => {
  let service: PlannerStateHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PlannerStateHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
