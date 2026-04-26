import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { NgOptimizedImage, CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormBuilder, Validators, FormsModule } from '@angular/forms';
import { HeroService, HeroSection } from '../../services/hero.service';
import { SocialService, SocialLink } from '../../services/social.service';
import { AuthService } from '../../services/auth.service';
import { MediaService } from '../../services/media.service';
import { SharedModalComponent } from '../shared/modal/modal';
import { ToastService } from '../../services/toast.service';

@Component({
  selector: 'app-hero',
  imports: [NgOptimizedImage, CommonModule, ReactiveFormsModule, FormsModule, SharedModalComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './hero.html',
  styleUrl: './hero.scss',
})
export class HeroComponent implements OnInit {
  private heroService = inject(HeroService);
  private socialService = inject(SocialService);
  private authService = inject(AuthService);
  private mediaService = inject(MediaService);
  private fb = inject(FormBuilder);
  private toastService = inject(ToastService);

  hero = this.heroService.hero;
  socialLinks = this.socialService.socialLinks;
  isAdmin = this.authService.isAdmin;
  
  isModalOpen = signal(false);
  isUploading = signal(false);

  heroForm = this.fb.group({
    headline: ['', Validators.required],
    subheadline: [''],
    biography: ['', Validators.required],
    profileImage: [''],
    resumeUrl: ['']
  });

  ngOnInit() {
    this.heroService.getHero().subscribe(data => {
      this.heroForm.patchValue(data);
    });
    this.socialService.getSocialLinks().subscribe();
  }

  openEditModal() {
    if (this.hero()) {
      this.heroForm.patchValue(this.hero()!);
    }
    this.isModalOpen.set(true);
  }

  closeModal() {
    this.isModalOpen.set(false);
  }

  onFileSelected(event: any, type: 'image' | 'resume') {
    const file = event.target.files[0];
    if (file) {
      this.isUploading.set(true);
      this.mediaService.uploadFile(file).subscribe({
        next: (res) => {
          if (type === 'image') {
            this.heroForm.patchValue({ profileImage: res.url });
          } else {
            this.heroForm.patchValue({ resumeUrl: res.url });
          }
          this.isUploading.set(false);
          this.toastService.show('Archivo subido con éxito', 'success');
        },
        error: () => {
          this.isUploading.set(false);
          this.toastService.show('Error al subir el archivo', 'error');
        }
      });
    }
  }

  saveHero() {
    if (this.heroForm.valid) {
      this.heroService.updateHero(this.heroForm.value as any).subscribe({
        next: () => {
          this.toastService.show('Hero actualizado', 'success');
          this.closeModal();
        },
        error: () => {
          this.toastService.show('Error al actualizar hero', 'error');
        }
      });
    }
  }

  updateSocial(link: SocialLink) {
    this.socialService.updateSocialLink(link.id, link).subscribe({
      next: () => this.toastService.show(`${link.platform} actualizado`, 'success'),
      error: () => this.toastService.show('Error al actualizar red social', 'error')
    });
  }
}
