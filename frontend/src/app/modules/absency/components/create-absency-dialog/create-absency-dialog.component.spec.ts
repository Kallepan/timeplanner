import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateAbsencyDialogComponent } from './create-absency-dialog.component';

describe('CreateAbsencyDialogComponent', () => {
  let component: CreateAbsencyDialogComponent;
  let fixture: ComponentFixture<CreateAbsencyDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateAbsencyDialogComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAbsencyDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
