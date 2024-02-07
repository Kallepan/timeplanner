import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonPreviewComponent } from './person-preview.component';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { Component } from '@angular/core';

@Component({
  imports: [PersonPreviewComponent],
  template: `<app-person-preview [persons]="persons">
    <div>T</div>
  </app-person-preview>`,
  standalone: true,
})
class TestComponent {
  persons: PersonWithMetadata[] = [];
}

describe('PersonPreviewComponent', () => {
  let component: PersonPreviewComponent;
  let fixture: ComponentFixture<PersonPreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PersonPreviewComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display person name', async () => {
    component.persons = [{ id: '1', last_name: 'Doe', first_name: 'John' } as PersonWithMetadata];
    fixture.detectChanges();

    expect(component.displayedPersonStrings).toEqual(['Doe (1)']);

    expect(fixture.nativeElement.textContent).toContain('Doe (1)');
  });

  it('should display empty card if no person', async () => {
    expect(component.displayedPersonStrings).toEqual([]);

    expect(fixture.nativeElement.textContent).toContain(' - ');
  });
});

describe('PersonPreviewComponent (TestComponent)', () => {
  let component: TestComponent;
  let fixture: ComponentFixture<TestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(TestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display the body content', () => {
    const body = fixture.nativeElement.querySelector('div');
    // The - is from the empty persons array
    expect(body.textContent).toBe(' - T');
  });

  it('should populate the persons', () => {
    component.persons = [{ id: '1', last_name: 'Doe', first_name: 'John' } as PersonWithMetadata];

    fixture.detectChanges();

    const body = fixture.nativeElement.querySelector('div');
    expect(body.textContent).toBe(' Doe (1) T');
  });
});
