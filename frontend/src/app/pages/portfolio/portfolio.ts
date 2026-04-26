import { Component, ChangeDetectionStrategy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HeaderComponent } from '../../components/header/header';
import { HeroComponent } from '../../components/hero/hero';
import { SkillsComponent } from '../../components/skills/skills';
import { ProjectsComponent } from '../../components/projects/projects';
import { ExperienceComponent } from '../../components/experience/experience';
import { EducationComponent } from '../../components/education/education';
import { FooterComponent } from '../../components/footer/footer';

@Component({
  selector: 'app-portfolio',
  imports: [
    CommonModule,
    HeaderComponent,
    HeroComponent,
    SkillsComponent,
    ProjectsComponent,
    ExperienceComponent,
    EducationComponent,
    FooterComponent
  ],
  templateUrl: './portfolio.html',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class PortfolioComponent {}
