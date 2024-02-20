import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfirmationDialogComponent, ConfirmationDialogComponentData } from './confirmation-dialog.component';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

describe('ConfirmationDialogComponent', () => {
  let component: ConfirmationDialogComponent;
  let fixture: ComponentFixture<ConfirmationDialogComponent>;

  let mockDialogData: ConfirmationDialogComponentData;

  beforeEach(async () => {
    mockDialogData = {
      title: 'Mock Title',
      confirmationMessage: 'Mock Confirmation Message',
    };

    await TestBed.configureTestingModule({
      imports: [ConfirmationDialogComponent],
      providers: [
        {
          provide: MAT_DIALOG_DATA,
          useValue: mockDialogData,
        },
        provideNoopAnimations(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfirmationDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize title', () => {
    expect(component.title).toEqual(mockDialogData.title);
  });

  it('should initialize confirmation message', () => {
    expect(component.confirmationMessage).toEqual(mockDialogData.confirmationMessage);
  });

  it('should initialize dummy control with true value', () => {
    expect(component.dummyControl.value).toEqual(true);
  });
});
