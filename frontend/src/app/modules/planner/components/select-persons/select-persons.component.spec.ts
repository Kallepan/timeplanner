import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SelectPersonsComponent } from './select-persons.component';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { provideNoopAnimations } from '@angular/platform-browser/animations';

describe('SelectPersonsComponent', () => {
  let component: SelectPersonsComponent;
  let fixture: ComponentFixture<SelectPersonsComponent>;
  let mockPersonDataContainerService: jasmine.SpyObj<PersonDataContainerService>;

  beforeEach(async () => {
    mockPersonDataContainerService = jasmine.createSpyObj('PersonDataContainerService', ['getPersons']);

    await TestBed.configureTestingModule({
      imports: [SelectPersonsComponent],
      providers: [{ provide: PersonDataContainerService, useValue: mockPersonDataContainerService }, provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(SelectPersonsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have initiallySetPersons', () => {
    expect(component.initiallySetPersons).toBeTruthy();
  });

  it('should emit commentEditRequest', () => {
    const button = fixture.nativeElement.querySelector('button#edit-comment-button');
    spyOn(component.commentEditRequest, 'emit');

    button.click();

    expect(component.commentEditRequest.emit).toHaveBeenCalled();
  });

  it('should not display delete button if comment is empty', () => {
    component.comment = '';
    expect(component.comment).toBe('');
    fixture.detectChanges();

    const button = fixture.nativeElement.querySelector('button#delete-comment-button');
    expect(button).toBeFalsy();
  });

  it('should emit commentDeleteRequest', () => {
    component.comment = 'Test comment';
    expect(component.comment).toBe('Test comment');
    fixture.detectChanges();

    const button = fixture.nativeElement.querySelector('button#delete-comment-button');
    expect(button).toBeTruthy();
    spyOn(component.commentDeleteRequest, 'emit');

    button.click();

    expect(component.commentDeleteRequest.emit).toHaveBeenCalled();
  });
});
