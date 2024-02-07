import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { PersonWithMetadata } from '@app/shared/interfaces/person';

@Component({
  selector: 'app-person-preview',
  standalone: true,
  imports: [CommonModule, MatCardModule],
  templateUrl: './person-preview.component.html',
  styleUrl: './person-preview.component.scss',
})
export class PersonPreviewComponent {
  displayedPersonStrings: string[] = [];
  @Input() set persons(value: PersonWithMetadata[]) {
    this.displayedPersonStrings = value.map((person) => `${person.last_name} (${person.id})`);
  }
}
