import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditableTimetableComponent } from './editable-timetable.component';

describe('EditableTimetableComponent', () => {
  let component: EditableTimetableComponent;
  let fixture: ComponentFixture<EditableTimetableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EditableTimetableComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(EditableTimetableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
