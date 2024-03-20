import { importProvidersFrom } from '@angular/core';
import { Routes } from '@angular/router';
import { HomeComponent } from './core/components/home/home.component';
import { ViewerModule } from './modules/viewer/viewer.module';
import { hasAccessToDepartmentGuard, isAdmin, isAuthenticated } from './core/guards/auth-guards';
import { PlannerModule } from './modules/planner/planner.module';
import { AbsencyModule } from './modules/absency/absency.module';

export const routes: Routes = [
  {
    path: 'planner',
    loadChildren: () => import('./modules/planner/planner.routes').then((m) => m.routes),
    canActivate: [isAuthenticated, hasAccessToDepartmentGuard],
    canActivateChild: [isAuthenticated, hasAccessToDepartmentGuard],
    providers: [importProvidersFrom(PlannerModule)],
  },
  {
    path: 'absency',
    loadChildren: () => import('./modules/absency/absency.routes').then((m) => m.routes),
    canActivate: [isAuthenticated, hasAccessToDepartmentGuard],
    canActivateChild: [isAuthenticated, hasAccessToDepartmentGuard],
    providers: [importProvidersFrom(AbsencyModule)],
  },
  {
    path: 'viewer',
    loadChildren: () => import('./modules/viewer/viewer.routes').then((m) => m.routes),
    providers: [importProvidersFrom(ViewerModule)],
  },
  {
    path: 'admin',
    loadChildren: () => import('./modules/admin/admin.routes').then((m) => m.routes),
    canActivate: [isAuthenticated, isAdmin],
  },
  { path: '', component: HomeComponent, title: 'Home' },
];
