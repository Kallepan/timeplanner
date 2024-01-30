import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SelectPersonComponent } from './select-person.component';

describe('SelectPersonComponent', () => {
  let component: SelectPersonComponent;
  let fixture: ComponentFixture<SelectPersonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SelectPersonComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(SelectPersonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
