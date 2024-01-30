import { importProvidersFrom } from '@angular/core';
import { Routes } from '@angular/router';
import { HomeComponent } from './core/components/home/home.component';
import { PlannerModule } from './modules/planner/planner.module';
import { ViewerModule } from './modules/viewer/viewer.module';
import { isAuthenticated } from './core/guards/auth-guard';

export const routes: Routes = [
  {
    path: 'planner',
    loadChildren: () => import('./modules/planner/planner.routes').then((m) => m.routes),
    canActivate: [isAuthenticated],
    canActivateChild: [isAuthenticated],
    providers: [importProvidersFrom(PlannerModule)],
  },
  {
    path: 'viewer',
    loadChildren: () => import('./modules/viewer/viewer.routes').then((m) => m.routes),
    providers: [importProvidersFrom(ViewerModule)],
  },
  { path: '', component: HomeComponent, title: 'Home' },
];
