import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SelectPersonForTimeslotComponent } from './select-person-for-timeslot.component';

describe('SelectPersonForTimeslotComponent', () => {
  let component: SelectPersonForTimeslotComponent;
  let fixture: ComponentFixture<SelectPersonForTimeslotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SelectPersonForTimeslotComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SelectPersonForTimeslotComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
