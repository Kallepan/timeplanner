import { DestroyRef, Injectable, inject } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatSnackBar, MatSnackBarHorizontalPosition, MatSnackBarVerticalPosition } from '@angular/material/snack-bar';
import { Subject, concatMap, map, of } from 'rxjs';

/**
 * This service is used to display notifications to the user.
 *
 */

type Message = {
  message: string;
  type: 'info' | 'warn';
};

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private _horizontalPosition: MatSnackBarHorizontalPosition = 'start';
  private _verticalPosition: MatSnackBarVerticalPosition = 'bottom';
  private _snackBar = inject(MatSnackBar);
  private _message = new Subject<Message>();

  private destroyRef$ = inject(DestroyRef);

  infoMessage(message: string) {
    this._message.next({ message, type: 'info' });
  }

  warnMessage(message: string) {
    this._message.next({ message, type: 'warn' });
  }

  private _getSnackBarDelay(message: Message) {
    const snackbarRef = this._snackBar._openedSnackBarRef;
    if (snackbarRef) {
      return snackbarRef.afterDismissed().pipe(map(() => message));
    }

    return of(message);
  }
  constructor() {
    this._message
      .pipe(
        concatMap((message) => this._getSnackBarDelay(message)),
        takeUntilDestroyed(this.destroyRef$),
      )
      .subscribe((res) => {
        this._snackBar.open(res.message, 'Dismiss', {
          duration: 2500,
          horizontalPosition: this._horizontalPosition,
          verticalPosition: this._verticalPosition,
          panelClass: `${res.type}-snackbar`,
        });
      });
  }
}
