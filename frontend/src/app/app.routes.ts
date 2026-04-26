import { Routes } from '@angular/router';
import { PortfolioComponent } from './pages/portfolio/portfolio';
import { LoginComponent } from './components/login/login';

export const routes: Routes = [
  { path: '', component: PortfolioComponent },
  { path: 'login', component: LoginComponent },
  { path: '**', redirectTo: '' }
];
