/**
 * This service is respsonsible for providing the data for the timetable.
 *
 * The data is provided from the backend as an array of Workday objects.
 * Each workday represents a single timeslot on a given day to which a person can be assigned to.
 *
 * These are grouped by workplace and then by the respective timeslots to generate the timetable.
 * Here we also populate the gridRowStart and gridRowEnd as well as the gridColumn (timeslot)
 * and gridRow (slotsGroup).
 *
 * The final data objects looks like this:
 * Workplace -> Has many TimeslotGroups -> of which each has many Timeslots.
 * In the html template we simple loop over the workplaces and then over the timeslotGroups and timeslots.
 **/
import { Injectable, computed, effect, inject, signal } from '@angular/core';
import { WorkdayTimeslot } from '@app/shared/interfaces/workday_timeslot';
import { catchError, from, map, mergeMap, of, reduce, tap } from 'rxjs';
import { DisplayedWorkplace } from '../../modules/viewer/interfaces/workplace';
import { convertWorkdaysToDisplayedWorkday } from '../functions/convert-workdays-to-displayed-workdays';
import { ActiveDepartmentHandlerService } from './active-department-handler.service';
import { ActiveWeekHandlerService } from './active-week-handler.service';
import { WorkdayAPIService } from './workday-api.service';

@Injectable({
  providedIn: null,
})
export class TimetableDataContainerService {
  protected _colorizeMissing = signal<boolean>(true);
  get colorizeMissing$(): boolean {
    return this._colorizeMissing();
  }
  set colorizeMissing(value: boolean) {
    this._colorizeMissing.set(value);
  }

  protected _isLoading = signal<boolean>(true);
  get isLoading$(): boolean {
    return this._isLoading();
  }

  protected _displayComments = signal<boolean>(true);
  get displayComments$(): boolean {
    return this._displayComments();
  }
  set displayComments(value: boolean) {
    this._displayComments.set(value);
  }

  protected _displayTimes = signal<boolean>(false);
  get displayTimes$(): boolean {
    return this._displayTimes();
  }
  set displayTimes(value: boolean) {
    this._displayTimes.set(value);
  }

  protected _colorize = signal<boolean>(true);
  get colorize$(): boolean {
    return this._colorize();
  }
  set colorize(value: boolean) {
    this._colorize.set(value);
  }

  addPersonWithWeekday(personId: string, weekday: number): void {
    // Theoretically we only weekdays 1 to 5, but store all here
    const presentWeekdays = this.mapOfPersonsAssignedToTheWholeWeek().get(personId) ?? new Set<number>();
    presentWeekdays.add(weekday);
    this.mapOfPersonsAssignedToTheWholeWeek.set(new Map(this.mapOfPersonsAssignedToTheWholeWeek().set(personId, presentWeekdays)));
  }
  removePersonWithWeekday(personId: string, weekday: number): void {
    // Theoretically we only weekdays 1 to 5, but store all here
    const presentWeekdays = this.mapOfPersonsAssignedToTheWholeWeek().get(personId);
    if (presentWeekdays) {
      presentWeekdays.delete(weekday);
      this.mapOfPersonsAssignedToTheWholeWeek.set(new Map(this.mapOfPersonsAssignedToTheWholeWeek().set(personId, presentWeekdays)));
    }
  }
  mapOfPersonsAssignedToTheWholeWeek = signal<Map<string, Set<number>>>(new Map<string, Set<number>>());
  listOfPersonsAssignedToTheWholeWeek = computed(() => {
    return Array.from(this.mapOfPersonsAssignedToTheWholeWeek().entries())
      .filter(([, presentWeekdays]) => presentWeekdays.has(1) && presentWeekdays.has(2) && presentWeekdays.has(3) && presentWeekdays.has(4) && presentWeekdays.has(5))
      .map(([personId]) => personId);
  });

  /**
   * This method is a setter for the `workdays` property. It takes an array of `WorkdayTimeslot` objects as input.
   * Each `WorkdayTimeslot` object represents a timeslot in a workday at a specific workplace.
   * The method organizes these timeslots into a structure that is convenient for displaying in a grid.
   *
   **/
  private _activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private _activeWeekHandlerService = inject(ActiveWeekHandlerService);
  private _workdayAPIService = inject(WorkdayAPIService);
  protected _workplaces = signal<DisplayedWorkplace[]>([]);
  get workplaces$(): DisplayedWorkplace[] {
    return this._workplaces() ?? [];
  }

  constructor() {
    effect(
      () => {
        from(this._activeWeekHandlerService.activeWeek$)
          .pipe(
            tap(() => this._isLoading.set(true)),
            map((weekday) => weekday.dateString),
            mergeMap((date) =>
              // fetch workdays for each date in the week
              this._workdayAPIService.getWorkdays(this._activeDepartmentHandlerService.activeDepartment$, date).pipe(
                map((resp) => resp.data), // map the response to the data property
                catchError(() => of([])), // if there is an error, return an empty array
              ),
            ),
            reduce((acc, workdays) => [...acc, ...workdays], [] as WorkdayTimeslot[]), // reduce the workdays into a single array
            map((workdays) => workdays.sort((a, b) => a.workplace.id.localeCompare(b.workplace.id))),
            map((workdays) => convertWorkdaysToDisplayedWorkday(workdays)),
            tap((workplaces) => {
              const presentWeekdaysPerPersonID = new Map<string, Set<number>>();

              workplaces.forEach((workplace) => {
                workplace.timeslotGroups.forEach((timeslotGroup) => {
                  timeslotGroup.workdayTimeslots.forEach((timeslot) => {
                    timeslot.persons.forEach((person) => {
                      const presentWeekdays = presentWeekdaysPerPersonID.get(person.id) ?? new Set<number>();
                      presentWeekdays.add(timeslot.weekday);
                      presentWeekdaysPerPersonID.set(person.id, presentWeekdays);
                    });
                  });
                });
              });

              this.mapOfPersonsAssignedToTheWholeWeek.set(presentWeekdaysPerPersonID);
            }),
          )
          .subscribe((workplaceGroups) => {
            this._isLoading.set(false);
            this._workplaces.set(workplaceGroups);
          });
      },
      { allowSignalWrites: true },
    );
  }
}
