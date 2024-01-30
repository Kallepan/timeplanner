import { Component, DestroyRef, OnInit, inject } from '@angular/core';
import { ViewerStateHandlerService } from '../../services/viewer-state-handler.service';
import { MatButtonModule } from '@angular/material/button';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, map } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { ViewOnlyTimetableComponent } from '@app/modules/viewer/components/view-only-timetable/view-only-timetable.component';
import { ActionsComponent } from '../actions/actions.component';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';

@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [MatButtonModule, RouterLink, AsyncPipe, ViewOnlyTimetableComponent, ActionsComponent],
  templateUrl: './landing-page.component.html',
  styleUrl: './landing-page.component.scss',
})
export class LandingPageComponent implements OnInit {
  timetableDataContainerService = inject(TimetableDataContainerService);
  // inject the services here
  viewerStateHandlerService = inject(ViewerStateHandlerService);
  private destroyRef$ = inject(DestroyRef);

  // router
  private route = inject(ActivatedRoute);

  ngOnInit(): void {
    // fetch department query param
    this.route.queryParams
      .pipe(
        takeUntilDestroyed(this.destroyRef$),
        // set the department
        map((params) => params['department']),
        filter((department): department is string => !!department),
        map((department) => department.toLowerCase()),
        // fetch the current date
        map((department) => {
          const currentDate = new Date();

          return {
            department,
            currentDate,
          };
        }),
      )
      .subscribe(({ department, currentDate }) => {
        // set both the department and the current date
        // This will cause the activeWeek signal to be updated
        // and fetch all workdays for the current week we want to view
        setTimeout(() => {
          this.viewerStateHandlerService.setActiveView(department, currentDate);
        }, 0);
      });
  }

  getEditRouterQueryParams(): { department?: string } {
    return {
      department: this.viewerStateHandlerService.activeWeek$?.department,
    };
  }
}
