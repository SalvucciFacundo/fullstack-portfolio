import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormBuilder, Validators } from '@angular/forms';
import { SkillService, Skill } from '../../services/skill.service';
import { AuthService } from '../../services/auth.service';
import { SharedModalComponent } from '../shared/modal/modal';
import { ToastService } from '../../services/toast.service';

@Component({
  selector: 'app-skills',
  imports: [CommonModule, ReactiveFormsModule, SharedModalComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './skills.html',
  styleUrl: './skills.scss',
})
export class SkillsComponent implements OnInit {
  private skillService = inject(SkillService);
  private authService = inject(AuthService);
  private fb = inject(FormBuilder);
  private toastService = inject(ToastService);

  skills = this.skillService.skills;
  isAdmin = this.authService.isAdmin;

  isModalOpen = signal(false);
  editingSkill = signal<Skill | null>(null);

  skillForm = this.fb.group({
    name: ['', Validators.required],
    iconClass: ['', Validators.required],
    category: ['frontend', Validators.required],
    displayOrder: [0]
  });

  ngOnInit() {
    this.skillService.getSkills().subscribe();
  }

  getSkillsByCategory(category: string) {
    return this.skills().filter(s => s.category === category);
  }

  openAddModal() {
    this.editingSkill.set(null);
    this.skillForm.reset({ category: 'frontend', displayOrder: 0 });
    this.isModalOpen.set(true);
  }

  openEditModal(skill: Skill) {
    this.editingSkill.set(skill);
    this.skillForm.patchValue(skill);
    this.isModalOpen.set(true);
  }

  closeModal() {
    this.isModalOpen.set(false);
  }

  saveSkill() {
    if (this.skillForm.valid) {
      const skillData = this.skillForm.value as Skill;
      const obs = this.editingSkill() 
        ? this.skillService.updateSkill(this.editingSkill()!.id!, skillData)
        : this.skillService.createSkill(skillData);

      obs.subscribe({
        next: () => {
          this.toastService.show(`Habilidad ${this.editingSkill() ? 'actualizada' : 'creada'}`, 'success');
          this.closeModal();
        },
        error: () => this.toastService.show('Error al guardar habilidad', 'error')
      });
    }
  }

  deleteSkill(id: string) {
    if (confirm('¿Estás seguro de eliminar esta habilidad?')) {
      this.skillService.deleteSkill(id).subscribe({
        next: () => this.toastService.show('Habilidad eliminada', 'success'),
        error: () => this.toastService.show('Error al eliminar', 'error')
      });
    }
  }
}
