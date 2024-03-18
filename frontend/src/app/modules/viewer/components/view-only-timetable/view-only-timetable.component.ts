/**
 * This component is used to display the timetable in a read-only mode.
 * It is used in the view-only-timetable module. To use it import this component,
 * inject the associated service and use it in your template. The service is used
 * to get format and handle the data from the backend.
 **/
import { Component, inject } from '@angular/core';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { TimetableDataContainerService } from '@app/shared/services/timetable-data-container.service';
import { ActionsComponent } from '../actions/actions.component';
import { PersonPreviewComponent } from '@app/shared/components/person-preview/person-preview.component';
import { MatButtonModule } from '@angular/material/button';
import { DatePipe, NgClass, NgStyle, TitleCasePipe } from '@angular/common';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-view-only-timetable',
  standalone: true,
  imports: [NgStyle, NgClass, DatePipe, TitleCasePipe, MatButtonModule, ActionsComponent, PersonPreviewComponent],
  templateUrl: './view-only-timetable.component.html',
  styleUrl: './view-only-timetable.component.scss',
})
export class ViewOnlyTimetableComponent {
  timetableDataContainerService = inject(TimetableDataContainerService);
  activeWeekdayService = inject(ActiveWeekHandlerService);
  themeHandlerService = inject(ThemeHandlerService);

  private httpClient = inject(HttpClient);
  print(element: HTMLDivElement) {
    let themeChanged = false;
    if (this.themeHandlerService.isDark$) {
      this.themeHandlerService.toggleTheme();
      themeChanged = true;
    }
    this.httpClient.get('assets/print-styles.css', { responseType: 'text' }).subscribe((styleString) => {
      const printWindow = window.open('', 'Druckansicht');
      printWindow?.document.write(`<html><head><style>${styleString}</style></head><body>${element.innerHTML}</body></html>`);
      printWindow?.document.close();
      printWindow?.focus();
      printWindow?.print();
      if (themeChanged) {
        this.themeHandlerService.toggleTheme();
      }
    });
  }
}
