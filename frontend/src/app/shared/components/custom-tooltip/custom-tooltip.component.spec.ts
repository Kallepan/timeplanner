import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomTooltipComponent } from './custom-tooltip.component';

describe('CustomTooltipComponent', () => {
  let component: CustomTooltipComponent;
  let fixture: ComponentFixture<CustomTooltipComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CustomTooltipComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(CustomTooltipComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have a text input', () => {
    expect(component.text).toBeDefined();

    component.text = 'test';
    fixture.detectChanges();

    expect(fixture.nativeElement.querySelector('.content').textContent).toBe(' test ');
  });

  it('should have a topCoordinate', () => {
    expect(component.topCoordinate).toBeDefined();
  });

  it('should have a leftCoordinate', () => {
    expect(component.leftCoordinate).toBeDefined();
  });

  it('should update topCoordinate and leftCoordinate on mousemove', () => {
    component.mouseX = 100;
    component.mouseY = 100;

    fixture.detectChanges();

    expect(component.topCoordinate).toBe('110px');
    expect(component.leftCoordinate).toBe('110px');
  });
});
