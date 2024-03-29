import { AsyncPipe, TitleCasePipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { FormControl } from '@angular/forms';
import { NavigationEnd, Router } from '@angular/router';
import { filter, map, switchMap } from 'rxjs';

@Component({
  selector: 'app-custom-header',
  standalone: true,
  imports: [AsyncPipe, TitleCasePipe],
  templateUrl: './custom-header.component.html',
  styleUrl: './custom-header.component.scss',
})
export class CustomHeaderComponent {
  private _router = inject(Router);

  control = new FormControl('');

  activatedRoute$ = this._router.events.pipe(
    filter((event) => event instanceof NavigationEnd),
    // Get the activated route
    map(() => this._router.routerState.root),
    // Get the last activated route
    map((route) => {
      while (route.firstChild) {
        route = route.firstChild;
      }
      return route;
    }),
    // Get the data from the route
    switchMap((route) => route.data),
    map((data) => data['featureFlag']),
  );
}
