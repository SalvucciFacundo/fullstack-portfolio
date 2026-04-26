import { Injectable, inject, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, tap } from 'rxjs';

export interface Skill {
  id?: string;
  name: string;
  iconClass: string;
  category: string;
  displayOrder: number;
}

@Injectable({
  providedIn: 'root'
})
export class SkillService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/skills`;
  private adminUrl = `${environment.apiUrl}/admin/skills`;

  skills = signal<Skill[]>([]);

  getSkills(): Observable<Skill[]> {
    return this.http.get<Skill[]>(this.apiUrl).pipe(
      tap(data => this.skills.set(data))
    );
  }

  createSkill(skill: Skill): Observable<Skill> {
    return this.http.post<Skill>(this.adminUrl, skill).pipe(
      tap(() => this.getSkills().subscribe())
    );
  }

  updateSkill(id: string, skill: Skill): Observable<Skill> {
    return this.http.put<Skill>(`${this.adminUrl}/${id}`, skill).pipe(
      tap(() => this.getSkills().subscribe())
    );
  }

  deleteSkill(id: string): Observable<void> {
    return this.http.delete<void>(`${this.adminUrl}/${id}`).pipe(
      tap(() => this.getSkills().subscribe())
    );
  }
}
