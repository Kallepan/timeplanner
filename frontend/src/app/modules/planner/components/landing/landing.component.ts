import { CommonModule } from '@angular/common';
import { Component, DestroyRef, OnInit, inject } from '@angular/core';
import { EditableTimetableComponent } from '../editable-timetable/editable-timetable.component';
import { ActionsComponent } from '../actions/actions.component';
import { PlannerStateHandlerService } from '../../services/planner-state-handler.service';
import { PersonListComponent } from '../person-list/person-list.component';
import { ActivatedRoute } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, map } from 'rxjs';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [CommonModule, EditableTimetableComponent, ActionsComponent, PersonListComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.scss',
  providers: [],
})
export class LandingComponent implements OnInit {
  timetableDataContainerService = inject(TimetableDataContainerService);
  plannerStateHandlerService = inject(PlannerStateHandlerService);
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
          this.plannerStateHandlerService.setActiveView(department, currentDate);
        }, 0);
      });
  }
}
