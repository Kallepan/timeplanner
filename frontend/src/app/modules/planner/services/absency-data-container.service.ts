import { Injectable, effect, inject, signal } from '@angular/core';
import { AbsenceByDate } from '@app/modules/absency/interfaces/absence';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { ActiveWeekHandlerService, Weekday } from '@app/shared/services/active-week-handler.service';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { catchError, forkJoin, map, of, switchMap, tap, throwError } from 'rxjs';

@Injectable({
  providedIn: null,
})
export class AbsencyDataContainerService {
  private activeWeekHandlerService = inject(ActiveWeekHandlerService);
  private activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private departmentAPIService = inject(DepartmentAPIService);

  private _absencesGroupedByWeekday = signal<{ weekday: Weekday; absences: AbsenceByDate[] }[]>([]);
  get absencesGroupedByWeekday$() {
    return this._absencesGroupedByWeekday();
  }

  private _loading = signal<boolean>(true);
  get loading$() {
    return this._loading();
  }

  constructor() {
    effect(
      () => {
        of(this.activeWeekHandlerService.activeWeek$)
          .pipe(
            tap(() => this._loading.set(true)),
            switchMap((weekDays) => {
              const obs = weekDays.map((weekDay) =>
                this.departmentAPIService.getAbsencesForDepartment(this.activeDepartmentHandlerService.activeDepartment$, weekDay.dateString).pipe(
                  catchError((err) => throwError(() => err)),
                  map((resp) => resp.data),
                  map((absences) => {
                    if (absences === null || absences === undefined) return [];

                    return absences.map((absence) => ({
                      personID: absence.person_id,
                      reason: absence.reason,
                      date: absence.date,
                      createdAt: absence.created_at,
                    }));
                  }),
                  map((absences) => ({ weekday: weekDay, absences })), // Pair the weekday with the absence data
                ),
              );

              return forkJoin(obs);
            }),
          )
          .subscribe((absences) => {
            setTimeout(() => {
              this._loading.set(false);
            }, 0);
            this._absencesGroupedByWeekday.set(absences);
          });
      },
      { allowSignalWrites: true },
    );
  }
}
