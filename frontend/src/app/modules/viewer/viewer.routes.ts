import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    title: 'Viewer',
    loadComponent: () => import('./components/viewer-landing-page/viewer-landing-page.component').then((m) => m.ViewerLandingPageComponent),
  },
];
