import { TestBed } from '@angular/core/testing';

import { PersonEditorDataContainerService } from './person-editor-data-container.service';

describe('PersonEditorDataContainerService', () => {
  let service: PersonEditorDataContainerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PersonEditorDataContainerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
