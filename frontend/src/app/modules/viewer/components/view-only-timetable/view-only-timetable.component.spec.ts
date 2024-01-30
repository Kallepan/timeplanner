import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewOnlyTimetableComponent } from './view-only-timetable.component';

describe('ViewOnlyTimetableComponent', () => {
  let component: ViewOnlyTimetableComponent;
  let fixture: ComponentFixture<ViewOnlyTimetableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ViewOnlyTimetableComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ViewOnlyTimetableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
