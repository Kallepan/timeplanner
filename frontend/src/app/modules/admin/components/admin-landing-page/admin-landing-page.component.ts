import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';

type RouteConfig = {
  path: string;
  name: string;
};

@Component({
  selector: 'app-admin-landing-page',
  standalone: true,
  imports: [CommonModule, MatButtonModule, RouterLink],
  templateUrl: './admin-landing-page.component.html',
  styleUrl: './admin-landing-page.component.scss',
})
export class AdminLandingPageComponent {
  routes: RouteConfig[] = [
    {
      path: 'schema',
      name: 'Schemaeditor',
    },
    {
      path: 'workday',
      name: 'Arbeitstagseditor',
    },
    {
      path: 'person',
      name: 'Personaleditor',
    },
  ];
}
