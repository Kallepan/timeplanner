import { TestBed } from '@angular/core/testing';

import { ViewerStateHandlerService } from './viewer-state-handler.service';
import { WorkdayAPIService } from '@app/shared/services/workday-api.service';

describe('ViewerStateHandlerService', () => {
  let service: ViewerStateHandlerService;
  let mockWorkdayAPIService: jasmine.SpyObj<WorkdayAPIService>;

  beforeEach(() => {
    mockWorkdayAPIService = jasmine.createSpyObj('WorkdayAPIService', ['getWorkdays']);

    TestBed.configureTestingModule({
      providers: [
        {
          provide: WorkdayAPIService,
          useValue: mockWorkdayAPIService,
        },
        ViewerStateHandlerService,
      ],
    });
    service = TestBed.inject(ViewerStateHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
