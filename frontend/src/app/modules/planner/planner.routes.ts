import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@app/modules/planner/components/planner-landing-page/planner-landing-page.component').then((m) => m.PlannerLandingPageComponent),
  },
];
