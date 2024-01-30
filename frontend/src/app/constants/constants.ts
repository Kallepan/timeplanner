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

  public static TITLE_SHORT = 'TP';
  public static TITLE_LONG = 'TP - Timeplanner';
  public static VERSION = '1.0.0';

  public static IS_PRODUCTION = environment.production;
}
