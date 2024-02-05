import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { DialogLayoutComponent } from './dialog-layout.component';

// Dummy Test component to host the layout and fill out the ng-content
@Component({
  imports: [DialogLayoutComponent],
  template: `
    <app-dialog-layout [title]="title" [control]="control">
      <div body>Test Content</div>
    </app-dialog-layout>
  `,
  standalone: true,
})
class TestComponent {
  title = 'Test Title';
  control = new FormControl('');
}

describe('DialogLayoutComponent', () => {
  let component: TestComponent;
  let fixture: ComponentFixture<TestComponent>;
  let mockDialogRef: jasmine.SpyObj<MatDialogRef<DialogLayoutComponent>>;

  beforeEach(async () => {
    mockDialogRef = jasmine.createSpyObj<MatDialogRef<DialogLayoutComponent>>('MatDialogRef', ['close']);

    await TestBed.configureTestingModule({
      imports: [TestComponent],
      providers: [
        {
          provide: MatDialogRef,
          useValue: mockDialogRef,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(TestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display the title', () => {
    const title = fixture.nativeElement.querySelector('.title');
    expect(title.textContent).toBe('Test Title');
  });

  it('should display the body content', () => {
    const body = fixture.nativeElement.querySelector('.body');
    expect(body.textContent).toBe('Test Content');
  });

  it('save-button should close the dialog with the new text', () => {
    // set control value
    component.control.setValue('New Text');
    fixture.detectChanges();

    // fetch button by id #save-button
    const button = fixture.nativeElement.querySelector('#save-button');

    // click button
    button.click();

    // assert dialog close
    expect(mockDialogRef.close).toHaveBeenCalledWith('New Text');
  });

  it('abort-button should close the dialog with null', () => {
    // fetch button by id #abort-button
    const button = fixture.nativeElement.querySelector('#abort-button');

    // click button
    button.click();

    // assert dialog close
    expect(mockDialogRef.close).toHaveBeenCalledWith(null);
  });

  it('save-button should be disabled if control is invalid', () => {
    // set control value
    component.control.setErrors({ required: true });
    fixture.detectChanges();

    // fetch button by id #save-button
    const button = fixture.nativeElement.querySelector('#save-button');

    // assert button is disabled
    expect(button.disabled).toBeTrue();
  });
});
