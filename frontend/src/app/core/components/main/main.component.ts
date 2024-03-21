import { OverlayContainer } from '@angular/cdk/overlay';
import { Component, HostBinding, effect, inject } from '@angular/core';
import { MatSidenavModule } from '@angular/material/sidenav';
import { RouterOutlet } from '@angular/router';
import { ThemeHandlerService } from '@app/core/services/theme-handler.service';
import { FooterComponent } from '../footer/footer.component';
import { HeaderComponent } from '../header/header.component';
import { SidenavComponent } from '../sidenav/sidenav.component';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss'],
  standalone: true,
  imports: [MatSidenavModule, HeaderComponent, FooterComponent, SidenavComponent, RouterOutlet],
})
export class MainComponent {
  // dependencies
  private _overlayContainer = inject(OverlayContainer);
  private _themeHandlerService = inject(ThemeHandlerService);

  // theme
  @HostBinding('class') get themeMode() {
    return this._themeHandlerService.isDark$ ? 'theme-dark' : 'theme-light';
  }

  get isDark() {
    return this._themeHandlerService.isDark$;
  }

  toggleTheme() {
    this._themeHandlerService.toggleTheme();
  }

  // lifecycle hooks
  constructor() {
    effect(() => {
      if (this._themeHandlerService.isDark$) {
        this._overlayContainer.getContainerElement().classList.add('theme-dark');
        this._overlayContainer.getContainerElement().classList.remove('theme-light');
      } else {
        this._overlayContainer.getContainerElement().classList.add('theme-light');
        this._overlayContainer.getContainerElement().classList.remove('theme-dark');
      }
    });
  }
}
