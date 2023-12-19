import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { MainComponent } from './core/components/main/main.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, MainComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {}
