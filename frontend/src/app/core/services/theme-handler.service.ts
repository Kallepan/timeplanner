import { Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ThemeHandlerService {
  protected _isDark = signal(true);
  get isDark$() {
    return this._isDark();
  }

  toggleTheme() {
    this._isDark.set(!this._isDark());
  }
}
