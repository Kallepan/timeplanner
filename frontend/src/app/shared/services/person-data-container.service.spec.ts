import { TestBed } from '@angular/core/testing';

import { PersonDataContainerService } from './person-data-container.service';

describe('PersonDataContainerService', () => {
  let service: PersonDataContainerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PersonDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
