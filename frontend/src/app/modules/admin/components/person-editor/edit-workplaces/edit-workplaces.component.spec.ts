import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditWorkplacesComponent } from './edit-workplaces.component';

describe('EditWorkplacesComponent', () => {
  let component: EditWorkplacesComponent;
  let fixture: ComponentFixture<EditWorkplacesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EditWorkplacesComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(EditWorkplacesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
