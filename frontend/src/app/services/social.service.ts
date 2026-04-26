import { Injectable, inject, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, tap } from 'rxjs';

export interface SocialLink {
  id: string;
  platform: string;
  url: string;
  iconName: string;
  isActive: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class SocialService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/social`;
  private adminUrl = `${environment.apiUrl}/admin/social`;

  socialLinks = signal<SocialLink[]>([]);

  getSocialLinks(): Observable<SocialLink[]> {
    return this.http.get<SocialLink[]>(this.apiUrl).pipe(
      tap(data => this.socialLinks.set(data))
    );
  }

  updateSocialLink(id: string, link: SocialLink): Observable<SocialLink> {
    return this.http.put<SocialLink>(`${this.adminUrl}/${id}`, link).pipe(
      tap(() => this.getSocialLinks().subscribe())
    );
  }
}
