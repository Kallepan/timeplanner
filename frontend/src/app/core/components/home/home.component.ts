import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { constants } from '@app/constants/constants';
import { RouteHandlerService } from '@app/core/services/route-handler.service';

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

  routeConfigurations = this.routeHandlerService.routeConfigurations();
}
