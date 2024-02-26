import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@app/modules/admin/components/admin-landing-page/admin-landing-page.component').then((m) => m.AdminLandingPageComponent),
  },
];
