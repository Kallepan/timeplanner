export class messages {
  public static AUTH = {
    LOGGED_IN: 'Login erfolgreich',
    LOGGED_OUT: 'Logout erfolgreich',

    INVALID_CREDENTIALS: 'Ungültiger Identifier oder Passwort',
    UNAUTHORIZED: 'Sie sind nicht berechtigt, diese Funktion auszuführen',
    FORBIDDEN: 'Sie dürfen diese Funktion nicht ausführen',
  };

  public static GENERAL = {
    BAD_REQUEST: 'Bad request',
    UNKNOWN_ERROR: 'An unknown error occured',
    UPDATE_FAILED: 'Update fehlgeschlagen',
    SERVER_ERROR: 'Ein Serverfehler ist aufgetreten',
    FEATURE_FLAG_DISABLED: 'Diese Funktion ist fuer Sie nicht verfügbar',
    FEATURE_NOT_IMPLEMENTED: 'Diese Funktion ist noch nicht implementiert',
    NO_RESULTS_FOUND: 'Keine Ergebnisse gefunden',

    FORM_ERRORS: {
      MIN: 'Wert muss größer oder gleich sein',
      MAX: 'Wert muss kleiner oder gleich sein',
      REQUIRED: 'Dieses Feld ist erforderlich',
      MIN_LENGTH: 'Mindestlänge ist',
      MAX_LENGTH: 'Maximale Länge ist',
      PATTERN: 'Ungültiges Format',
    },
  };
}
