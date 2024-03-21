import { Component, computed, inject } from '@angular/core';
import { MatChipSelectionChange, MatChipsModule } from '@angular/material/chips';
import { messages } from '@app/core/constants/messages';
import { NotificationService } from '@app/core/services/notification.service';
import { PersonEditorDataContainerService } from '@app/modules/admin/services/person-editor-data-container.service';
import { PersonAPIService } from '@app/shared/services/person-api.service';
import { filter, of, switchMap } from 'rxjs';

@Component({
  selector: 'app-edit-weekdays',
  standalone: true,
  imports: [MatChipsModule],
  templateUrl: './edit-weekdays.component.html',
  styleUrl: './edit-weekdays.component.scss',
})
export class EditWeekdaysComponent {
  private _personEditorDataContainerService = inject(PersonEditorDataContainerService);
  private _personAPIService = inject(PersonAPIService);
  private _notificationService = inject(NotificationService);

  weekdays = computed(() => {
    return this._personEditorDataContainerService.weekdays$.map((weekday) => ({
      ...weekday,
      selected: this._personEditorDataContainerService.activePerson$?.weekdays?.some((selectedWeekday) => selectedWeekday.id === weekday.id) ?? false,
    }));
  });

  selectWeekday(event: MatChipSelectionChange) {
    of(event)
      .pipe(
        filter((event) => event.isUserInput),
        filter(() => !!this._personEditorDataContainerService.activePerson$),
        switchMap((event) => {
          if (event.selected) {
            return this._personAPIService.addWeekdayToPerson(event.source.value, this._personEditorDataContainerService.activePerson$!.id);
          }

          return this._personAPIService.removeWeekdayFromPerson(event.source.value, this._personEditorDataContainerService.activePerson$!.id);
        }),
      )
      .subscribe({
        next: () => {
          this._notificationService.infoMessage(messages.ADMIN.PERSON_WEEKDAY_UPDATED);
        },
        error: () => {
          this._notificationService.warnMessage(messages.ADMIN.PERSON_WEEKDAY_UPDATE_FAILED);
          event.source.toggleSelected();
        },
      });
  }
}
