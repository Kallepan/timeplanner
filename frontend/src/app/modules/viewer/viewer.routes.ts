import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    title: 'Viewer',
    loadComponent: () =>
      import('./components/landing-page/landing-page.component').then(
        (m) => m.LandingPageComponent,
      ),
  },
];
