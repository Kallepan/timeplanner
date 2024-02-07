import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonPreviewComponent } from './person-preview.component';
import { PersonWithMetadata } from '@app/shared/interfaces/person';

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

    expect(fixture.nativeElement.textContent).toContain('John Doe');
  });

  it('should display empty card if no person', async () => {
    expect(component.persons).toEqual([]);

    expect(fixture.nativeElement.textContent).toContain('No person found');
  });
});
