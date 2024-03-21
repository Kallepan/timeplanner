export class messages {
  public static ADMIN = {
    CREATE_SUCCESSFUL: 'Erfolgreich erstellt',
    CREATE_FAILED: 'Hinzuügen fehlgeschlagen',
    DELETE_SUCCESSFUL: 'Erfolgreich gelöscht',
    DELETE_FAILED: 'Löschen fehlgeschlagen',

    TIMESLOT_WEEKDAY_UNASSIGNED: 'Wochentag erfolgreich entfernt',
    TIMESLOT_WEEKDAY_UNASSIGNMENT_FAILED: 'Entfernen des Wochentags fehlgeschlagen',
    TIMESLOT_WEEKDAY_ASSIGNED: 'Wochentag erfolgreich zugeordnet',
    TIMESLOT_WEEKDAY_ASSIGNMENT_FAILED: 'Zuordnen des Wochentags fehlgeschlagen',
    TIMESLOT_WEEKDAY_UPDATE_FAILED: 'Aktualisieren des Wochentags fehlgeschlagen',
    TIMESLOT_WEEKDAY_UPDATE_SUCCESS: 'Wochentag erfolgreich aktualisiert',

    PERSON_WEEKDAY_UPDATED: 'Wochentage der Person erfolgreich aktualisiert',
    PERSON_WEEKDAY_UPDATE_FAILED: 'Aktualisieren der Wochentage der Person fehlgeschlagen',
    PERSON_WORKPLACE_UPDATED: 'Arbeitsplätze der Person erfolgreich aktualisiert',
    PERSON_WORKPLACE_UPDATE_FAILED: 'Aktualisieren der Arbeitsplätze der Person fehlgeschlagen',
  };
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

  public static ABSENCY = {
    ALREADY_EXISTS: 'Es existiert bereits ein Abwesenheitsgrund an diesem Tag',
    CREATED: 'Abwesenheit erfolgreich erstellt',
    DELETED: 'Abwesenheit erfolgreich gelöscht',

    DELETE_CONFIRMATION: 'Sind Sie sicher, dass Sie die Abwesenheit löschen möchten?',
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
