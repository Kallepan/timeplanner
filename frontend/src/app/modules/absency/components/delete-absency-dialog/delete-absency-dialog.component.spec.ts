import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeleteAbsencyDialogComponent, DeleteAbsencyDialogComponentData } from './delete-absency-dialog.component';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { provideNoopAnimations } from '@angular/platform-browser/animations';

describe('DeleteAbsencyDialogComponent', () => {
  let component: DeleteAbsencyDialogComponent;
  let fixture: ComponentFixture<DeleteAbsencyDialogComponent>;
  let mockDialogData: DeleteAbsencyDialogComponentData;

  beforeEach(async () => {
    mockDialogData = {
      personID: 'Mock Person ID',
      date: new Date(),
    };
    await TestBed.configureTestingModule({
      imports: [DeleteAbsencyDialogComponent],
      providers: [{ provide: MAT_DIALOG_DATA, useValue: mockDialogData }, provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(DeleteAbsencyDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize title with person ID', () => {
    expect(component.title).toEqual(`Abwesenheit löschen für ${mockDialogData.personID}`);
  });

  it('should initialize date', () => {
    expect(component.date).toEqual(mockDialogData.date);
  });

  it('should initialize dummy control with true value', () => {
    expect(component.dummyControl.value).toEqual(true);
  });
});
