import { ComponentFixture, TestBed } from '@angular/core/testing';

import { YearCalendarComponent } from './year-calendar.component';
import { By } from '@angular/platform-browser';
import { ActivePersonHandlerServiceService } from '../../services/active-person-handler-service.service';

describe('YearCalendarComponent', () => {
  let component: YearCalendarComponent;
  let fixture: ComponentFixture<YearCalendarComponent>;
  let mockActivePersonHandlerServiceService: jasmine.SpyObj<ActivePersonHandlerServiceService>;

  beforeEach(async () => {
    mockActivePersonHandlerServiceService = jasmine.createSpyObj('ActivePersonHandlerServiceService', ['handleDayClick']);
    await TestBed.configureTestingModule({
      providers: [{ provide: ActivePersonHandlerServiceService, useValue: mockActivePersonHandlerServiceService }],
      imports: [YearCalendarComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(YearCalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have a calendar', () => {
    const compiled = fixture.debugElement.query(By.css('#calendar'));
    expect(compiled).toBeTruthy();
  });

  it('should update the tooltip text', () => {
    const compiled = fixture.debugElement.query(By.css('.tooltip'));
    expect(compiled).toBeFalsy();

    component.tooltipText = 'Test';
    component.tooltipVisible = true;
    fixture.detectChanges();

    const compiledAfterUpdate = fixture.debugElement.query(By.css('.tooltip'));
    expect(compiledAfterUpdate).toBeTruthy();
    expect(compiledAfterUpdate.nativeElement).toBeTruthy();
    expect(compiledAfterUpdate.nativeElement.textContent).toContain('Test');
  });
});
