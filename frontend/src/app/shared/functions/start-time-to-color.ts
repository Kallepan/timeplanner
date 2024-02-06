export const startTimeToColorForLightMode = (startTime: string): string | null => {
  switch (startTime) {
    case '06:00':
      return '#FFC0C0';
    case '07:00':
      return '#FFFF99';
    case '07:45':
      return '#C0FFC0';
    case '10:45':
      return '#C5D9F1';
    default:
      return null;
  }
};

export const startTimeToColorForDarkMode = (startTime: string): string | null => {
  switch (startTime) {
    case '06:00':
      return '#851209';
    case '07:00':
      return '#857209';
    case '07:45':
      return '#09850d';
    case '10:45':
      return '#0d0985';
    default:
      return null;
  }
};
