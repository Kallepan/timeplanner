import { TestBed } from '@angular/core/testing';

import { ThemeHandlerService } from './theme-handler.service';

describe('ThemeHandlerService', () => {
  let service: ThemeHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ThemeHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('isDark$ should be true', () => {
    expect(service.isDark$).toBeTrue();
  });

  it('isDark$ should be false after toggleTheme', () => {
    service.toggleTheme();
    expect(service.isDark$).toBeFalse();
  });
});
