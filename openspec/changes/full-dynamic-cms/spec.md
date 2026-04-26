# Change Specification: Full Dynamic CMS

## 1. Goal
Transform the static portfolio into a fully dynamic CMS where the admin can manage all content from the UI.

## 2. Requirements

### 2.1. Hero Section
- **Dynamic Content**: Headline, subheadline, and biography must be editable.
- **Profile Image**: Admin must be able to upload a new profile image.
- **Resume (CV)**: Admin must be able to upload a new PDF for the resume.

### 2.2. Social Links
- **Social Media**: Manage GitHub, LinkedIn, and other social links.
- **Attributes**: URL, platform name, and icon (Material Symbols or DevIcons).
- **Design**: Icons must be larger and more prominent.

### 2.3. Skills & Tools
- **Categories**: Frontend, Backend, QA, Tools.
- **CRUD**: Ability to add, edit, and delete skills.
- **Attributes**: Name, Icon class.

### 2.4. Professional Experience
- **CRUD**: Full management of work history.
- **Attributes**: Company, Role, Description (rich text or bullet points), Start Date, End Date (optional if current).
- **UI**: Consistent "Add" button and modal-based entry.

### 2.5. Education
- **CRUD**: Full management of educational background.
- **Attributes**: Institution, Degree, Start Date, End Date.
- **UI**: Consistent "Add" button and modal-based entry.

### 2.6. Media Management
- **Storage**: Local filesystem storage in `backend/uploads/` directory.
- **Endpoints**: `POST /api/upload` to handle image/PDF uploads.
- **Security**: Only authenticated admin can upload files.

## 3. User Experience (UX)
- All "Add" and "Edit" actions must open a **Modal** with a reactive form.
- Buttons should have consistent styling across sections.
- Real-time updates in the UI after saving (using Angular Signals).

## 4. Success Criteria
- [ ] Backend provides all dynamic content via APIs.
- [ ] Admin can update any section without editing code.
- [ ] Images and CV files can be replaced from the UI.
- [ ] The "Authenticate" logic is fully integrated with protected routes.
