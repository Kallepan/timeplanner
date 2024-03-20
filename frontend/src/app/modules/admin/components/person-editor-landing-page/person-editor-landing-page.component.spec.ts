import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonEditorLandingPageComponent } from './person-editor-landing-page.component';

describe('PersonEditorLandingPageComponent', () => {
  let component: PersonEditorLandingPageComponent;
  let fixture: ComponentFixture<PersonEditorLandingPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PersonEditorLandingPageComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonEditorLandingPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
