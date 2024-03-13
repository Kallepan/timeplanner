import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WorkdayEditorComponent } from './workday-editor.component';

describe('WorkdayEditorComponent', () => {
  let component: WorkdayEditorComponent;
  let fixture: ComponentFixture<WorkdayEditorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WorkdayEditorComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(WorkdayEditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
