import { Injectable, inject, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, tap } from 'rxjs';

export interface Education {
  id?: string;
  institution: string;
  degree: string;
  startDate: string;
  endDate?: string;
}

@Injectable({
  providedIn: 'root'
})
export class EducationService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/education`;
  private adminUrl = `${environment.apiUrl}/admin/education`;

  education = signal<Education[]>([]);

  getEducation(): Observable<Education[]> {
    return this.http.get<Education[]>(this.apiUrl).pipe(
      tap(data => this.education.set(data))
    );
  }

  createEducation(edu: Education): Observable<Education> {
    return this.http.post<Education>(this.adminUrl, edu).pipe(
      tap(() => this.getEducation().subscribe())
    );
  }

  updateEducation(id: string, edu: Education): Observable<Education> {
    return this.http.put<Education>(`${this.adminUrl}/${id}`, edu).pipe(
      tap(() => this.getEducation().subscribe())
    );
  }

  deleteEducation(id: string): Observable<void> {
    return this.http.delete<void>(`${this.adminUrl}/${id}`).pipe(
      tap(() => this.getEducation().subscribe())
    );
  }
}
