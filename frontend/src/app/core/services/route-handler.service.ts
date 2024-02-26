import { Injectable, signal } from '@angular/core';

type RouteConfiguration = {
  path: string;
  title: string;
  id?: string;
};

@Injectable({
  providedIn: 'root',
})
export class RouteHandlerService {
  // route configurations for the viewer
  routeConfigurations = signal<RouteConfiguration[]>([
    {
      path: '/viewer',
      title: 'MIBI - Planung',
      id: 'bak',
    },
    {
      path: '/absency',
      title: 'MIBI - Abwesenheit',
      id: 'bak',
    },
    {
      path: '/admin',
      title: 'Admin',
      id: '',
    },
  ]);
}
