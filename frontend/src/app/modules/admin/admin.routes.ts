import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@app/modules/admin/components/admin-landing-page/admin-landing-page.component').then((m) => m.AdminLandingPageComponent),
  },
  {
    path: 'schema',
    loadComponent: () => import('@app/modules/admin/components/schema-editor/schema-editor.component').then((m) => m.SchemaEditorComponent),
  },
  {
    path: 'workday',
    loadComponent: () => import('@app/modules/admin/components/workday-editor/workday-editor.component').then((m) => m.WorkdayEditorComponent),
  },
  {
    path: 'person',
    loadComponent: () => import('@app/modules/admin/components/person-editor/person-editor.component').then((m) => m.PersonEditorComponent),
  },
];
