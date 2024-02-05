import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { PersonWithMetadata } from '@app/shared/interfaces/person';

@Component({
  selector: 'app-edit-person-preview',
  standalone: true,
  imports: [CommonModule, MatCardModule],
  templateUrl: './edit-person-preview.component.html',
  styleUrl: './edit-person-preview.component.scss',
})
export class EditPersonPreviewComponent {
  @Input() person: PersonWithMetadata | null = null;
  @Input({ required: true }) comment: string;
}
