import { TestBed } from '@angular/core/testing';

import { PlannerTimetableDataService } from './planner-timetable-data.service';

describe('PlannerTimetableDataService', () => {
  let service: PlannerTimetableDataService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PlannerTimetableDataService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
