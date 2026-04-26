import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';

export interface Project {
  id?: number;
  title: string;
  description: string;
  image: string;
  category: string;
  techStack: string[];
  liveUrl?: string;
  githubUrl?: string;
  statusLabel?: string;
}

@Injectable({
  providedIn: 'root'
})
export class ProjectService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/projects`;

  // Público
  getAll(): Observable<Project[]> {
    return this.http.get<Project[]>(this.apiUrl);
  }

  // Admin
  create(project: Project): Observable<Project> {
    return this.http.post<Project>(`${environment.apiUrl}/admin/projects`, project);
  }

  update(id: number, project: Project): Observable<Project> {
    return this.http.put<Project>(`${environment.apiUrl}/admin/projects/${id}`, project);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${environment.apiUrl}/admin/projects/${id}`);
  }
}
