import { Component, inject } from '@angular/core';
import { ViewerStateHandlerService } from '../../services/viewer-state-handler.service';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [MatButtonModule, RouterLink],
  templateUrl: './landing-page.component.html',
  styleUrl: './landing-page.component.scss',
})
export class LandingPageComponent {
  // inject the services here
  private viewerStateHandlerService = inject(ViewerStateHandlerService);
}
