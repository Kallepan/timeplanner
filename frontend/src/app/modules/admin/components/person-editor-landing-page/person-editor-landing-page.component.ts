import { Component, OnInit, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { RouterLink } from '@angular/router';
import { DepartmentWithMetadata } from '@app/shared/interfaces/department';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { map, tap } from 'rxjs';

@Component({
  selector: 'app-person-editor-landing-page',
  standalone: true,
  imports: [MatProgressSpinnerModule, RouterLink, MatButtonModule],
  templateUrl: './person-editor-landing-page.component.html',
  styleUrl: './person-editor-landing-page.component.scss',
})
export class PersonEditorLandingPageComponent implements OnInit {
  departmentAPIService = inject(DepartmentAPIService);

  private _departments = signal<DepartmentWithMetadata[]>([]);
  get departments$() {
    return this._departments();
  }
  isLoading: boolean = true;
  ngOnInit(): void {
    this.departmentAPIService
      .getDepartments()
      .pipe(
        map((resp) => resp.data),
        tap(() => (this.isLoading = false)),
      )
      .subscribe({
        next: (departments) => {
          this._departments.set(departments);
        },
      });
  }
}
