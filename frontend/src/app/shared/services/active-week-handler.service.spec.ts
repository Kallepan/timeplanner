import { TestBed } from '@angular/core/testing';

import { ActiveWeekHandlerService } from './active-week-handler.service';

describe('ActiveWeekHandlerService', () => {
  let service: ActiveWeekHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ActiveWeekHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
