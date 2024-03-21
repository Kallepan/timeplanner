import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonEditorComponent } from './person-editor.component';

describe('PersonEditorComponent', () => {
  let component: PersonEditorComponent;
  let fixture: ComponentFixture<PersonEditorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PersonEditorComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonEditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
