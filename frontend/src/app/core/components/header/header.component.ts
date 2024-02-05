import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { constants } from '@app/constants/constants';
import { CustomHeaderComponent } from '../custom-header/custom-header.component';
import { LoginComponent } from '../login/login.component';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss'],
  standalone: true,
  imports: [CommonModule, MatSlideToggleModule, MatTooltipModule, MatToolbarModule, MatIconModule, MatButtonModule, LoginComponent, CustomHeaderComponent],
})
export class HeaderComponent {
  title = constants.TITLE_SHORT;
  @Input() isDark = false;
  @Output() themeToggled = new EventEmitter<void>();
  @Output() sidenavToggled = new EventEmitter<void>();
}
