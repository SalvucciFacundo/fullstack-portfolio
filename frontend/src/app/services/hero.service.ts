import { Injectable, inject, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, tap } from 'rxjs';

export interface HeroSection {
  id: string;
  headline: string;
  subheadline: string;
  biography: string;
  profileImage: string;
  resumeUrl: string;
  updatedAt: string;
}

@Injectable({
  providedIn: 'root'
})
export class HeroService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/hero`;
  private adminUrl = `${environment.apiUrl}/admin/hero`;

  hero = signal<HeroSection | null>(null);

  getHero(): Observable<HeroSection> {
    return this.http.get<HeroSection>(this.apiUrl).pipe(
      tap(data => this.hero.set(data))
    );
  }

  updateHero(hero: Partial<HeroSection>): Observable<HeroSection> {
    return this.http.put<HeroSection>(this.adminUrl, hero).pipe(
      tap(data => this.hero.set(data))
    );
  }
}
