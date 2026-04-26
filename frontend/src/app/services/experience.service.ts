import { Injectable, inject, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, tap } from 'rxjs';

export interface Experience {
  id?: string;
  company: string;
  role: string;
  description: string;
  startDate: string;
  endDate?: string;
  isCurrent: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class ExperienceService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/experience`;
  private adminUrl = `${environment.apiUrl}/admin/experience`;

  experiences = signal<Experience[]>([]);

  getExperiences(): Observable<Experience[]> {
    return this.http.get<Experience[]>(this.apiUrl).pipe(
      tap(data => this.experiences.set(data))
    );
  }

  createExperience(exp: Experience): Observable<Experience> {
    return this.http.post<Experience>(this.adminUrl, exp).pipe(
      tap(() => this.getExperiences().subscribe())
    );
  }

  updateExperience(id: string, exp: Experience): Observable<Experience> {
    return this.http.put<Experience>(`${this.adminUrl}/${id}`, exp).pipe(
      tap(() => this.getExperiences().subscribe())
    );
  }

  deleteExperience(id: string): Observable<void> {
    return this.http.delete<void>(`${this.adminUrl}/${id}`).pipe(
      tap(() => this.getExperiences().subscribe())
    );
  }
}
