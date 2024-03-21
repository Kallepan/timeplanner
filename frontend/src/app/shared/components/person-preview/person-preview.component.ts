import { Component, Input } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { PersonWithMetadata } from '@app/shared/interfaces/person';

@Component({
  selector: 'app-person-preview',
  standalone: true,
  imports: [MatCardModule],
  templateUrl: './person-preview.component.html',
  styleUrl: './person-preview.component.scss',
})
export class PersonPreviewComponent {
  displayedPersonString: string = '';
  @Input() largeText: boolean = false;
  @Input() set persons(value: PersonWithMetadata[]) {
    this.displayedPersonString = value.map((person) => `${person.last_name} (${person.id})`).join(', ');
  }
}
