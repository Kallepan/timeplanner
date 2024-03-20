import { Component } from '@angular/core';
import { PersonAutocompleteComponent } from '@app/modules/absency/components/person-autocomplete/person-autocomplete.component';

@Component({
  selector: 'app-person-editor',
  standalone: true,
  imports: [PersonAutocompleteComponent],
  templateUrl: './person-editor.component.html',
  styleUrl: './person-editor.component.scss',
})
export class PersonEditorComponent {}
