import { OverlayContainer } from '@angular/cdk/overlay';
import { CommonModule } from '@angular/common';
import { Component, HostBinding, effect, inject, signal } from '@angular/core';
import { MatSidenavModule } from '@angular/material/sidenav';
import { RouterOutlet } from '@angular/router';
import { FooterComponent } from '../footer/footer.component';
import { HeaderComponent } from '../header/header.component';
import { SidenavComponent } from '../sidenav/sidenav.component';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    MatSidenavModule,
    HeaderComponent,
    FooterComponent,
    SidenavComponent,
    RouterOutlet,
  ],
})
export class MainComponent {
  // dependencies
  private _overlayContainer = inject(OverlayContainer);

  // theme
  private _isDark = signal(true);
  @HostBinding('class') get themeMode() {
    return this._isDark() ? 'theme-dark' : 'theme-light';
  }

  get isDark() {
    return this._isDark();
  }

  toggleTheme() {
    this._isDark.set(!this._isDark());
  }

  // lifecycle hooks
  constructor() {
    effect(() => {
      if (this._isDark()) {
        this._overlayContainer.getContainerElement().classList.add('theme-dark');
        this._overlayContainer.getContainerElement().classList.remove('theme-light');
      } else {
        this._overlayContainer.getContainerElement().classList.add('theme-light');
        this._overlayContainer.getContainerElement().classList.remove('theme-dark');
      }
    });
  }
}
