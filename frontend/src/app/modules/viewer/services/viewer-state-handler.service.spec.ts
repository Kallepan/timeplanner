import { TestBed } from '@angular/core/testing';

import { ViewerStateHandlerService } from './viewer-state-handler.service';

describe('ViewerStateHandlerService', () => {
  let service: ViewerStateHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ViewerStateHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
