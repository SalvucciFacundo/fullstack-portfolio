# Technical Design: Full Dynamic CMS

## 1. Database Schema (PostgreSQL)

### Table: `hero_section`
- `id`: UUID (PK)
- `headline`: TEXT
- `subheadline`: TEXT
- `biography`: TEXT
- `profile_image`: TEXT (path)
- `resume_url`: TEXT (path)
- `updated_at`: TIMESTAMP

### Table: `skills`
- `id`: UUID (PK)
- `name`: TEXT
- `icon_class`: TEXT
- `category`: TEXT (frontend, backend, qa, tools)
- `display_order`: INTEGER

### Table: `social_links`
- `id`: UUID (PK)
- `platform`: TEXT
- `url`: TEXT
- `icon_name`: TEXT
- `is_active`: BOOLEAN

### Table: `experience`
- `id`: UUID (PK)
- `company`: TEXT
- `role`: TEXT
- `description`: TEXT
- `start_date`: DATE
- `end_date`: DATE (Nullable)
- `is_current`: BOOLEAN

### Table: `education`
- `id`: UUID (PK)
- `institution`: TEXT
- `degree`: TEXT
- `start_date`: DATE
- `end_date`: DATE

## 2. API Design (Go Backend)

### Base Endpoints
- `GET /api/portfolio/hero` -> Returns hero data.
- `PUT /api/admin/hero` -> Updates hero data (Auth Required).
- `GET /api/portfolio/skills` -> Returns skills list.
- `POST/PUT/DELETE /api/admin/skills` -> Skill management.
- `GET /api/portfolio/experience` -> Returns experience list.
- `POST/PUT/DELETE /api/admin/experience` -> Experience management.
- `GET /api/portfolio/education` -> Returns education list.
- `POST/PUT/DELETE /api/admin/education` -> Education management.

### Media Upload
- `POST /api/admin/upload` -> Multipart form-data for images/PDFs.
- **Storage**: Files saved to `backend/uploads/` with UUID-based names to avoid collisions.

## 3. Frontend Architecture (Angular)

### Services
- `HeroService`: Signal-based state for hero data.
- `SkillService`: CRUD operations for skills.
- `ExperienceService`: CRUD for experiences.
- `EducationService`: CRUD for education.
- `MediaService`: Handles file upload logic.

### UI Components
- **UnifiedModalComponent**: A reusable base component for admin modals.
- **FormGroup**: Using `ReactiveFormsModule` for all entries.
- **Image Preview**: Show current image vs new image before upload.

## 4. Security
- All `POST/PUT/DELETE` endpoints strictly checked against `authInterceptor` (JWT).
- Backend validates token in `AuthMiddleware`.
