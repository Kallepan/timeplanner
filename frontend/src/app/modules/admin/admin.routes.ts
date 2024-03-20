import { Routes } from '@angular/router';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';

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
    loadComponent: () => import('@app/modules/admin/components/person-editor-landing-page/person-editor-landing-page.component').then((m) => m.PersonEditorLandingPageComponent),
    providers: [DepartmentAPIService],
  },
  {
    path: 'person/detail',
    loadComponent: () => import('@app/modules/admin/components/person-editor/person-editor.component').then((m) => m.PersonEditorComponent),
    providers: [PersonDataContainerService, ActiveDepartmentHandlerService],
  },
];
