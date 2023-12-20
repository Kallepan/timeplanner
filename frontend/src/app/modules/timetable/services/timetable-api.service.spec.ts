import { TestBed } from '@angular/core/testing';

import { TimetableAPIService } from './timetable-api.service';

describe('TimetableAPIService', () => {
  let service: TimetableAPIService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TimetableAPIService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
