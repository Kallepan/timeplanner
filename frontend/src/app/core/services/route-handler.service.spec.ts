import { TestBed } from '@angular/core/testing';

import { RouteHandlerService } from './route-handler.service';

describe('RouteHandlerService', () => {
  let service: RouteHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RouteHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
