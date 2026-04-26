# Especificaciones: Backend en Go con Arquitectura Hexagonal (SQL Puro)

## REST API Endpoints

| Método | Endpoint | Descripción | Auth |
| --- | --- | --- | --- |
| GET | `/health` | Health check del sistema | No |
| POST | `/auth/login` | Login de admin (devuelve JWT) | No |
| GET | `/profile` | Obtiene info del Hero y About Me | No |
| PUT | `/profile` | Actualiza info del perfil | **SÍ (Admin)** |
| GET | `/projects` | Lista todos los proyectos | No |
| POST | `/projects` | Crea un proyecto | **SÍ (Admin)** |
| PUT | `/projects/:id` | Actualiza un proyecto | **SÍ (Admin)** |
| DELETE | `/projects/:id` | Borra un proyecto | **SÍ (Admin)** |
| GET | `/experience` | Lista experiencias laborales | No |
| POST | `/experience` | Crea una experiencia | **SÍ (Admin)** |
| PUT | `/experience/:id` | Actualiza una experiencia | **SÍ (Admin)** |
| DELETE | `/experience/:id` | Borra una experiencia | **SÍ (Admin)** |
| GET | `/education` | Lista formación académica | No |
| POST | `/education` | Crea una formación | **SÍ (Admin)** |
| PUT | `/education/:id` | Actualiza una formación | **SÍ (Admin)** |
| DELETE | `/education/:id` | Borra una formación | **SÍ (Admin)** |

## Modelos de Dominio

### User
- `ID`: UUID (Primary Key)
- `Email`: String (Unique, Required)
- `PasswordHash`: String (Required)
- `Role`: String (Default: 'admin')
- `CreatedAt`: Timestamp

### Profile (Hero/About Me)
- `ID`: UUID (Primary Key)
- `Name`: String
- `HeroTitle`: String
- `HeroSubtitle`: String
- `AboutMe`: Text
- `AvatarURL`: String
- `ResumeURL`: String
- `SocialLinks`: JSONB

### Project
- `ID`: UUID (Primary Key)
- `Title`: String (Required)
- `Description`: Text (Required)
- `ImageURL`: String
- `TechStack`: Array of Strings (Postgres `text[]`)
- `GithubURL`: String
- `LiveURL`: String
- `Category`: String
- `CreatedAt`: Timestamp

### Experience
- `ID`: UUID (Primary Key)
- `Company`: String
- `Position`: String
- `Period`: String
- `Description`: Text
- `TechStack`: Array of Strings
- `CreatedAt`: Timestamp

### Education
- `ID`: UUID (Primary Key)
- `Institution`: String
- `Degree`: String
- `Period`: String
- `Description`: Text
- `CreatedAt`: Timestamp

## Estrategia de Persistencia (SQL Puro)

- **Librería**: `jmoiron/sqlx` para el mapeo de structs.
- **Driver**: `jackc/pgx/v5`.
- **Migraciones**: Archivos `.sql` numerados en `migrations/`. 

## User Experience (Frontend Integration)

### Sistema de Notificaciones (Toasts)
- El frontend implementará un `NotificationService` basado en **Signals**.
- **Success**: Mensajes de éxito en Login, Create, Update y Delete.
- **Error**: Manejo global de errores HTTP para mostrar mensajes descriptivos del backend.
- **Feedback visual**: Spinners en botones durante las peticiones asíncronas.

## Reglas de Negocio y Validación

### Seguridad
- JWT para administración.
- Contraseñas hasheadas con `bcrypt`.

## Manejo de Errores
- Respuestas estándar: `{"message": "string", "status": int}`.
- El backend debe devolver mensajes claros para que el Toast del front sea útil.
