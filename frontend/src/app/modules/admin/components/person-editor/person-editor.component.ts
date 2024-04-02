import { Component, inject } from '@angular/core';
import { PersonAutocompleteComponent } from '@app/modules/absency/components/person-autocomplete/person-autocomplete.component';
import { PersonEditorDataContainerService } from '../../services/person-editor-data-container.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { EditWorkplacesComponent } from './edit-workplaces/edit-workplaces.component';
import { EditWeekdaysComponent } from './edit-weekdays/edit-weekdays.component';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { filter, map, switchMap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { PersonEditDialogComponent, PersonEditDialogComponentData } from './person-edit-dialog/person-edit-dialog.component';

@Component({
  selector: 'app-person-editor',
  standalone: true,
  imports: [PersonAutocompleteComponent, MatProgressSpinnerModule, EditWorkplacesComponent, EditWeekdaysComponent, MatButtonModule],
  templateUrl: './person-editor.component.html',
  styleUrl: './person-editor.component.scss',
  providers: [PersonEditorDataContainerService],
})
export class PersonEditorComponent {
  private _personEditorDataContainerService = inject(PersonEditorDataContainerService);
  private _personAPIService = inject(PersonAPIService);
  private _matDialog = inject(MatDialog);

  setActivePerson(personID: string) {
    this._personAPIService
      .getPerson(personID)
      .pipe(map((resp) => resp.data))
      .subscribe((person) => {
        this._personEditorDataContainerService.activePerson = person;
      });
  }
  getActivePerson() {
    return this._personEditorDataContainerService.activePerson$;
  }

  editPerson(personID: string) {
    this._personAPIService
      .getPerson(personID)
      .pipe(
        map((resp) => resp.data),
        switchMap((data) => {
          const dialogData = {
            first_name: data.first_name,
            last_name: data.last_name,
            email: data.email,
            active: data.active,
            working_hours: data.working_hours,
          };

          const matDialogConfig = new MatDialogConfig<PersonEditDialogComponentData>();
          matDialogConfig.enterAnimationDuration = 300;
          matDialogConfig.exitAnimationDuration = 300;
          matDialogConfig.disableClose = true;
          matDialogConfig.data = dialogData;

          return this._matDialog.open(PersonEditDialogComponent, matDialogConfig).afterClosed();
        }),
        filter((data) => !!data),
      )
      .subscribe((person) => {
        console.log('Editing person', person);
      });
  }
}
