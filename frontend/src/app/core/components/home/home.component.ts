import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { constants } from '@app/constants/constants';
import { AuthService } from '@app/core/services/auth.service';
import { RouteHandlerService } from '@app/core/services/route-handler.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, MatButtonModule, RouterLink],
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
})
export class HomeComponent {
  title = constants.TITLE_LONG;
  private routeHandlerService = inject(RouteHandlerService);
  private activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private authService = inject(AuthService);

  isAdmin() {
    return this.authService.isAdmin$;
  }
  setActiveDepartment(department: string | undefined) {
    this.activeDepartmentHandlerService.activeDepartment = department;
  }

  routeConfigurations = this.routeHandlerService.routeConfigurations();
}
