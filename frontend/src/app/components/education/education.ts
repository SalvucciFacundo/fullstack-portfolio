import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormBuilder, Validators } from '@angular/forms';
import { EducationService, Education } from '../../services/education.service';
import { AuthService } from '../../services/auth.service';
import { SharedModalComponent } from '../shared/modal/modal';
import { ToastService } from '../../services/toast.service';

@Component({
  selector: 'app-education',
  imports: [CommonModule, ReactiveFormsModule, SharedModalComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './education.html',
  styleUrl: './education.scss',
})
export class EducationComponent implements OnInit {
  private educationService = inject(EducationService);
  private authService = inject(AuthService);
  private fb = inject(FormBuilder);
  private toastService = inject(ToastService);

  education = this.educationService.education;
  isAdmin = this.authService.isAdmin;

  isModalOpen = signal(false);
  editingEducation = signal<Education | null>(null);

  educationForm = this.fb.group({
    institution: ['', Validators.required],
    degree: ['', Validators.required],
    startDate: ['', Validators.required],
    endDate: ['']
  });

  ngOnInit() {
    this.educationService.getEducation().subscribe();
  }

  openAddModal() {
    this.editingEducation.set(null);
    this.educationForm.reset();
    this.isModalOpen.set(true);
  }

  openEditModal(edu: Education) {
    this.editingEducation.set(edu);
    const formValue = { ...edu };
    if (edu.startDate) formValue.startDate = new Date(edu.startDate).toISOString().split('T')[0];
    if (edu.endDate) formValue.endDate = new Date(edu.endDate).toISOString().split('T')[0];
    
    this.educationForm.patchValue(formValue as any);
    this.isModalOpen.set(true);
  }

  closeModal() {
    this.isModalOpen.set(false);
  }

  saveEducation() {
    if (this.educationForm.valid) {
      const data = this.educationForm.value as Education;
      const obs = this.editingEducation()
        ? this.educationService.updateEducation(this.editingEducation()!.id!, data)
        : this.educationService.createEducation(data);

      obs.subscribe({
        next: () => {
          this.toastService.show(`Educación ${this.editingEducation() ? 'actualizada' : 'creada'}`, 'success');
          this.closeModal();
        },
        error: () => this.toastService.show('Error al guardar', 'error')
      });
    }
  }

  deleteEducation(id: string) {
    if (confirm('¿Eliminar esta entrada de educación?')) {
      this.educationService.deleteEducation(id).subscribe({
        next: () => this.toastService.show('Educación eliminada', 'success'),
        error: () => this.toastService.show('Error al eliminar', 'error')
      });
    }
  }
}
