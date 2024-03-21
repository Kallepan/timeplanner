import { Component, computed, inject } from '@angular/core';
import { MatChipSelectionChange, MatChipsModule } from '@angular/material/chips';
import { messages } from '@app/core/constants/messages';
import { NotificationService } from '@app/core/services/notification.service';
import { PersonEditorDataContainerService } from '@app/modules/admin/services/person-editor-data-container.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { filter, of, switchMap } from 'rxjs';

@Component({
  selector: 'app-edit-workplaces',
  standalone: true,
  imports: [MatChipsModule],
  templateUrl: './edit-workplaces.component.html',
  styleUrl: './edit-workplaces.component.scss',
})
export class EditWorkplacesComponent {
  private _personEditorDataContainerService = inject(PersonEditorDataContainerService);
  private _activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private _personAPIService = inject(PersonAPIService);
  private _notificationService = inject(NotificationService);

  workplaces = computed(() => {
    return this._personEditorDataContainerService.workplaces$.map((workplace) => ({
      ...workplace,
      selected:
        this._personEditorDataContainerService.activePerson$?.workplaces?.some(
          (selectedWorkplace) => selectedWorkplace.id === workplace.id && selectedWorkplace.department_id === workplace.department_id,
        ) ?? false,
    }));
  });

  selectWorkplace(event: MatChipSelectionChange) {
    of(event)
      .pipe(
        filter((event) => event.isUserInput),
        filter(() => !!this._personEditorDataContainerService.activePerson$),
        filter(() => !!this._activeDepartmentHandlerService.activeDepartment$),
        switchMap((event) => {
          if (event.selected) {
            return this._personAPIService.addWorkplaceToPerson(this._activeDepartmentHandlerService.activeDepartment$!, event.source.value, this._personEditorDataContainerService.activePerson$!.id);
          }

          return this._personAPIService.removeWorkplaceFromPerson(
            this._activeDepartmentHandlerService.activeDepartment$!,
            event.source.value,
            this._personEditorDataContainerService.activePerson$!.id,
          );
        }),
      )
      .subscribe({
        next: () => {
          this._notificationService.infoMessage(messages.ADMIN.PERSON_WORKPLACE_UPDATED);
        },
        error: () => {
          this._notificationService.warnMessage(messages.ADMIN.PERSON_WORKPLACE_UPDATE_FAILED);
          event.source.toggleSelected();
        },
      });
  }
}
