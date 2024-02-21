import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@app/modules/absency/components/absency-landing-page/absency-landing-page.component').then((m) => m.AbsencyLandingPageComponent),
  },
];
