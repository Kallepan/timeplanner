import { ComponentFixture, TestBed } from '@angular/core/testing';

import { YearCalendarComponent } from './year-calendar.component';
import { By } from '@angular/platform-browser';

describe('YearCalendarComponent', () => {
  let component: YearCalendarComponent;
  let fixture: ComponentFixture<YearCalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
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
});
