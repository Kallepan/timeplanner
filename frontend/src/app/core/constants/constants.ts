import { ValidWeekday } from '@app/core/types/weekday';
import { environment } from '@env/environment';

export class constants {
  public static APIS = {
    AUTH: environment.authUrl,
    BASE: environment.apiUrl,
    PLANNER: `${environment.apiUrl}/planner`,
  };

  public static JWT = {
    // One hour time in minutes, seconds and miliseconds
    ACCESS_STORAGE: 'access_token',
    REFRESH_STORAGE: 'refresh_token',
    USERNAME_STORAGE: 'username',
  };

  public static DEFAULT_TIMEOUT = 10000;

  public static TITLE_SHORT = 'DPP';
  public static TITLE_LONG = 'DPP - Dienstplan Planer';
  public static VERSION = '1.0.0';

  public static IS_PRODUCTION = environment.production;

  public static ABSENCY_REASONS = new Map<string, string>([
    ['Urlaub', '#2C8FC9'],
    ['Krank (AU)', '#9CB703'],
    ['Krank (keine AU)', '#F5BB00'],
    ['Fortbildung', '#B56CE2'],
  ]);

  public static POSSIBLE_WEEKDAYS: ValidWeekday[] = [
    { id: 1, name: 'Montag' },
    { id: 2, name: 'Dienstag' },
    { id: 3, name: 'Mittwoch' },
    { id: 4, name: 'Donnerstag' },
    { id: 5, name: 'Freitag' },
    { id: 6, name: 'Samstag' },
    { id: 7, name: 'Sonntag' },
  ];
}
