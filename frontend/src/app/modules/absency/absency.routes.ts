import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@app/modules/absency/components/landing/landing.component').then((m) => m.LandingComponent),
  },
];
