import { Injectable, signal, computed, inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { tap, catchError } from 'rxjs/operators';
import { of, Observable } from 'rxjs';
import { ToastService } from './toast.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private http = inject(HttpClient);
  private toast = inject(ToastService);
  private platformId = inject(PLATFORM_ID);
  
  // Signal para el token
  private tokenSignal = signal<string | null>(
    isPlatformBrowser(this.platformId) ? localStorage.getItem('token') : null
  );
  
  // Computed para el estado de login
  isAdmin = computed(() => !!this.tokenSignal());

  login(email: string, password: string): Observable<any> {
    return this.http.post(`${environment.apiUrl}/auth/login`, { email, password }).pipe(
      tap((res: any) => {
        if (res.token) {
          if (isPlatformBrowser(this.platformId)) {
            localStorage.setItem('token', res.token);
          }
          this.tokenSignal.set(res.token);
          this.toast.success('Welcome back, Admin!');
        }
      })
    );
  }

  logout() {
    if (isPlatformBrowser(this.platformId)) {
      localStorage.removeItem('token');
    }
    this.tokenSignal.set(null);
    this.toast.info('Logged out successfully');
  }

  getToken() {
    return this.tokenSignal();
  }
}
