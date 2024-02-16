import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AbsencyPanelComponent } from './absency-panel.component';

describe('AbsencyPanelComponent', () => {
  let component: AbsencyPanelComponent;
  let fixture: ComponentFixture<AbsencyPanelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AbsencyPanelComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AbsencyPanelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
