export class messages {
  public static AUTH = {
    LOGGED_IN: 'Login erfolgreich',
    LOGGED_OUT: 'Logout erfolgreich',
    LOGIN_FAILED: 'Login fehlgeschlagen',

    INVALID_CREDENTIALS: 'Ungültiger Identifier oder Passwort',
    UNAUTHORIZED: 'Sie sind nicht berechtigt, diese Funktion auszuführen',
    FORBIDDEN: 'Sie dürfen diese Funktion nicht ausführen',
  };

  public static GENERAL = {
    HTTP_ERROR: {
      BAD_REQUEST: 'Bad request',
      UNKNOWN_ERROR: 'An unknown error occured',
      UPDATE_FAILED: 'Update fehlgeschlagen',
      NOT_FOUND: 'Nicht gefunden',
      SERVER_ERROR: 'Serverfehler',
    },

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

  public static PLANNER = {
    TIMESLOT_ASSIGNMENT: {
      PERSON_ALREADY_ASSIGNED: 'Person ist bereits an einem anderen Zeitpunkt zugewiesen',
      PERSON_NOT_WORKING: 'Person arbeitet regulär an diesem Tag nicht',
      PERSON_NOT_QUALIFIED: 'Person ist nicht für diesen Arbeitsplatz qualifiziert',
      PERSON_ABSENT: 'Person ist an diesem Tag abwesend (Krank, Urlaub, etc.)',
      SUCCESS: 'Person erfolgreich zugeordnet',
    },

    TIMESLOT_UNASSIGNMENT: {
      SUCCESS: 'Person erfolgreich entfernt',
    },
  };
}
