import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditPersonPreviewComponent } from './edit-person-preview.component';

describe('EditPersonPreviewComponent', () => {
  let component: EditPersonPreviewComponent;
  let fixture: ComponentFixture<EditPersonPreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EditPersonPreviewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(EditPersonPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
