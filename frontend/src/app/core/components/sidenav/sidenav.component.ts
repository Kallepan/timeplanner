import { Component, EventEmitter, Output, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { AuthService } from '@app/core/services/auth.service';
import { RouteHandlerService } from '@app/core/services/route-handler.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
@Component({
  selector: 'app-sidenav',
  standalone: true,
  imports: [MatButtonModule, RouterLink],
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss'],
})
export class SidenavComponent {
  @Output() closeSidenav = new EventEmitter<void>();

  private routeHandlerService = inject(RouteHandlerService);
  private activeDepartmentHandlerService = inject(ActiveDepartmentHandlerService);
  private authService = inject(AuthService);

  isAdmin() {
    return this.authService.isAdmin$;
  }
  setActiveDepartment(department?: string) {
    if (!department) return;
    this.activeDepartmentHandlerService.activeDepartment = department;
  }

  routeConfigurations = this.routeHandlerService.routeConfigurations();
}
