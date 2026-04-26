import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormBuilder, Validators } from '@angular/forms';
import { ExperienceService, Experience } from '../../services/experience.service';
import { AuthService } from '../../services/auth.service';
import { SharedModalComponent } from '../shared/modal/modal';
import { ToastService } from '../../services/toast.service';

@Component({
  selector: 'app-experience',
  imports: [CommonModule, ReactiveFormsModule, SharedModalComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './experience.html',
  styleUrl: './experience.scss',
})
export class ExperienceComponent implements OnInit {
  private experienceService = inject(ExperienceService);
  private authService = inject(AuthService);
  private fb = inject(FormBuilder);
  private toastService = inject(ToastService);

  experiences = this.experienceService.experiences;
  isAdmin = this.authService.isAdmin;

  isModalOpen = signal(false);
  editingExperience = signal<Experience | null>(null);

  experienceForm = this.fb.group({
    company: ['', Validators.required],
    role: ['', Validators.required],
    description: ['', Validators.required],
    startDate: ['', Validators.required],
    endDate: [''],
    isCurrent: [false]
  });

  ngOnInit() {
    this.experienceService.getExperiences().subscribe();
  }

  openAddModal() {
    this.editingExperience.set(null);
    this.experienceForm.reset({ isCurrent: false });
    this.isModalOpen.set(true);
  }

  openEditModal(exp: Experience) {
    this.editingExperience.set(exp);
    // Convert dates to string for input[type="date"]
    const formValue = { ...exp };
    if (exp.startDate) formValue.startDate = new Date(exp.startDate).toISOString().split('T')[0];
    if (exp.endDate) formValue.endDate = new Date(exp.endDate).toISOString().split('T')[0];
    
    this.experienceForm.patchValue(formValue as any);
    this.isModalOpen.set(true);
  }

  closeModal() {
    this.isModalOpen.set(false);
  }

  saveExperience() {
    if (this.experienceForm.valid) {
      const data = this.experienceForm.value as Experience;
      const obs = this.editingExperience()
        ? this.experienceService.updateExperience(this.editingExperience()!.id!, data)
        : this.experienceService.createExperience(data);

      obs.subscribe({
        next: () => {
          this.toastService.show(`Experiencia ${this.editingExperience() ? 'actualizada' : 'creada'}`, 'success');
          this.closeModal();
        },
        error: () => this.toastService.show('Error al guardar', 'error')
      });
    }
  }

  deleteExperience(id: string) {
    if (confirm('¿Eliminar esta experiencia?')) {
      this.experienceService.deleteExperience(id).subscribe({
        next: () => this.toastService.show('Experiencia eliminada', 'success'),
        error: () => this.toastService.show('Error al eliminar', 'error')
      });
    }
  }
}
