import { importProvidersFrom } from '@angular/core';
import { Routes } from '@angular/router';
import { HomeComponent } from './core/components/home/home.component';
import { ViewerModule } from './modules/viewer/viewer.module';
import { hasAccessToDepartmentGuard, isAuthenticated } from './core/guards/auth-guard';
import { PlannerModule } from './modules/planner/planner.module';

export const routes: Routes = [
  {
    path: 'planner',
    loadChildren: () => import('./modules/planner/planner.routes').then((m) => m.routes),
    canActivate: [isAuthenticated, hasAccessToDepartmentGuard],
    canActivateChild: [isAuthenticated, hasAccessToDepartmentGuard],
    providers: [importProvidersFrom(PlannerModule)],
  },
  {
    path: 'viewer',
    loadChildren: () => import('./modules/viewer/viewer.routes').then((m) => m.routes),
    providers: [importProvidersFrom(ViewerModule)],
  },
  { path: '', component: HomeComponent, title: 'Home' },
];
