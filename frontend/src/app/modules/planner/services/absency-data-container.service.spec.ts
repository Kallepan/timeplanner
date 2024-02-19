import { TestBed } from '@angular/core/testing';

import { AbsencyDataContainerService } from './absency-data-container.service';

describe('AbsencyDataContainerService', () => {
  let service: AbsencyDataContainerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AbsencyDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
