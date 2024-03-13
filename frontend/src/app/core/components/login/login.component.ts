import { CommonModule } from '@angular/common';
import { Component, type OnInit, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { AuthService } from '@app/core/services/auth.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, MatProgressBarModule, MatIconModule, MatButtonModule, MatMenuModule, MatFormFieldModule, MatInputModule, ReactiveFormsModule, MatProgressSpinnerModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  authService = inject(AuthService);
  private readonly _fb = inject(FormBuilder);

  loginForm = this._fb.group({
    identifier: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(5)]],
    password: ['', [Validators.required]],
  });

  ngOnInit(): void {
    this.authService.verifyToken();
  }

  onSubmitLogin(): void {
    this.authService.login(this.loginForm.controls.identifier.value, this.loginForm.controls.password.value);
    this.loginForm.reset();
  }

  isLoading() {
    return this.authService.loading$;
  }
}
