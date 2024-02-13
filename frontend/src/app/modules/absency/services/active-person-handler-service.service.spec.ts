import { TestBed } from '@angular/core/testing';

import { ActivePersonHandlerServiceService } from './active-person-handler-service.service';

describe('ActivePersonHandlerServiceService', () => {
  let service: ActivePersonHandlerServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ActivePersonHandlerServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
