# Implementation Tasks: Full Dynamic CMS

## Phase 1: Backend Infrastructure
- [x] Create SQL migrations for `hero_section`, `skills`, `social_links`, `experience`, `education`.
- [x] Implement Go Domain models and Repository interfaces.
- [x] Implement Database repositories for all new entities.
- [x] Create API Handlers and register routes in `router.go`.
- [x] Implement Media Upload service (save to disk).

## Phase 2: Frontend Foundation
- [x] Create `HeroService`, `SkillService`, `ExperienceService`, `EducationService`, `MediaService`.
- [x] Update `ProjectService` if any cleanup is needed.
- [x] Create `SharedModalComponent` for reuse across admin actions.

## Phase 3: Hero & Social Refactor
- [x] Refactor `HeroComponent` to fetch data from backend.
- [x] Create `EditHeroModal` for headline/bio/image updates.
- [x] Update social links styling and make them dynamic.

## Phase 4: Skills, Experience & Education
- [x] Refactor `SkillsComponent` to be dynamic.
- [x] Refactor `ExperienceComponent` to fetch from DB.
- [x] Refactor `EducationComponent` to fetch from DB.
- [x] Implement "Add/Edit" modals for each section.

## Phase 5: Polish & Deployment Prep
- [ ] Ensure all admin actions are protected by `isAdmin` signal.
- [ ] Fix styling of all buttons to be consistent.
- [ ] Final verification of upload functionality on local server.
