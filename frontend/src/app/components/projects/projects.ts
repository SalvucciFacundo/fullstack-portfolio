import { ChangeDetectionStrategy, Component, viewChild, ElementRef, inject, signal, OnInit } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { ProjectService, Project } from '../../services/project.service';
import { AuthService } from '../../services/auth.service';
import { ToastService } from '../../services/toast.service';

@Component({
  selector: 'app-projects',
  imports: [CommonModule, NgOptimizedImage],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './projects.html',
  styleUrl: './projects.scss',
})
export class ProjectsComponent implements OnInit {
  private projectService = inject(ProjectService);
  private authService = inject(AuthService);
  private toast = inject(ToastService);
  
  slider = viewChild<ElementRef<HTMLDivElement>>('slider');
  projects = signal<Project[]>([]);
  isAdmin = this.authService.isAdmin;

  ngOnInit() {
    this.loadProjects();
  }

  loadProjects() {
    this.projectService.getAll().subscribe(data => {
      this.projects.set(data || []);
    });
  }

  scroll(direction: 'left' | 'right') {
    const sliderEl = this.slider()?.nativeElement;
    if (!sliderEl) return;

    const scrollAmount = 412;
    const newScrollPosition = direction === 'left' 
      ? sliderEl.scrollLeft - scrollAmount 
      : sliderEl.scrollLeft + scrollAmount;

    sliderEl.scrollTo({
      left: newScrollPosition,
      behavior: 'smooth'
    });
  }

  onAdd() {
    this.toast.info('Add project modal coming soon!');
  }

  onEdit(project: Project) {
    this.toast.info(`Editing ${project.title}`);
  }

  onDelete(project: Project) {
    if (confirm(`Delete project "${project.title}"?`)) {
      this.projectService.delete(project.id!).subscribe(() => {
        this.toast.success('Project deleted successfully');
        this.loadProjects();
      });
    }
  }
}
