import { TestBed } from '@angular/core/testing';

import { TimetableDataContainerService } from './timetable-data-container.service';

describe('TimetableDataContainerService', () => {
  let service: TimetableDataContainerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TimetableDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
