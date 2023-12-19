import { environment } from '@env/environment';

export class constants {
  public static APIS = {
    AUTH: environment.authUrl,
    BASE: environment.apiUrl,
    BAK: {
      BASE: environment.apiUrl + '/bak',
    },
    PCR: {
      BASE: environment.apiUrl + '/pcr',
    },
  };

  public static JWT = {
    // One hour time in minutes, seconds and miliseconds
    ACCESS_STORAGE: 'access_token',
    REFRESH_STORAGE: 'refresh_token',
    USERNAME_STORAGE: 'username',
  };

  public static TITLE_SHORT = 'ZMT';
  public static TITLE_LONG = 'ZMT - Zeitmanagement Tool';
  public static VERSION = '1.0.0';

  public static ROUTES = [{ path: '', title: 'Home', tooltip: 'Home' }];

  public static IS_PRODUCTION = environment.production;
}
