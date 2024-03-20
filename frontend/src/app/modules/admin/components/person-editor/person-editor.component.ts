import { Component, inject } from '@angular/core';
import { PersonAutocompleteComponent } from '@app/modules/absency/components/person-autocomplete/person-autocomplete.component';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { PersonEditorDataContainerService } from '../../services/person-editor-data-container.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

@Component({
  selector: 'app-person-editor',
  standalone: true,
  imports: [PersonAutocompleteComponent, MatProgressSpinnerModule],
  templateUrl: './person-editor.component.html',
  styleUrl: './person-editor.component.scss',
  providers: [PersonEditorDataContainerService],
})
export class PersonEditorComponent {
  private _personEditorDataContainerService = inject(PersonEditorDataContainerService);

  setActivePerson(person: PersonWithMetadata) {
    this._personEditorDataContainerService.activePerson = person;
  }
  getActivePerson() {
    return this._personEditorDataContainerService.activePerson$;
  }
}
