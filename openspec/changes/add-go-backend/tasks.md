# Plan de Batalla: Backend en Go (Hexagonal + SQL Puro)

Este plan detalla las tareas para implementar el backend del portfolio siguiendo Arquitectura Hexagonal, usando SQL puro (`sqlx` + `pgx`), `Echo v4` para el servidor y `samber/do` para la inyección de dependencias.

## Batch 1: Setup del Proyecto (Los Cimientos)
- [x] **Inicializar el módulo de Go**: Ejecutar `go mod init portfolio-backend` en la carpeta `backend/`.
- [x] **Estructura Hexagonal**: Crear el árbol de directorios:
    - `cmd/api`
    - `internal/core/domain`
    - `internal/core/ports`
    - `internal/core/services`
    - `internal/adapters/handlers`
    - `internal/adapters/repositories`
    - `internal/adapters/auth`
    - `internal/config`
    - `pkg/db`
    - `migrations`
- [x] **Docker Compose**: Configurar un archivo `docker-compose.yaml` con Postgres 15+ y pgAdmin.
- [x] **Makefile**: Crear comandos para automatizar el desarrollo (`run`, `test`, `migrate-up`, `migrate-down`).

## Batch 2: Infraestructura y Seguridad (El Motor)
- [x] **Conexión a DB**: Implementar el pool de conexiones en `pkg/db` usando `jackc/pgx/v5`.
- [x] **Migraciones**: Configurar `golang-migrate` y crear los archivos `.sql` iniciales para `users`, `profiles`, `projects`, `experiences` y `education`.
- [x] **Auth Service**: Implementar hashing de passwords con `bcrypt` y generación/validación de JWT.
- [x] **Middleware**: Crear un middleware de Echo para validar el token JWT y proteger las rutas de escritura.

## Batch 3: Dominio y Persistencia (El Corazón)
- [x] **Entidades de Dominio**: Definir los structs de Go en `internal/core/domain` (User, Profile, Project, Experience, Education).
- [x] **Puertos (Interfaces)**: Definir las interfaces `Repository` y `Service` en `internal/core/ports`.
- [x] **Repositorios Postgres**: Implementar las interfaces usando SQL puro con `sqlx` en `internal/adapters/repositories`.
- [x] **Seed Inicial**: Crear una migración para insertar el usuario Admin inicial.

## Batch 4: Aplicación y Handlers (Los Cables)
- [x] **Servicios de Aplicación**: Implementar la lógica de negocio en `internal/core/services`.
- [x] **Handlers de Echo**: Crear los controladores HTTP en `internal/adapters/handlers` para cada recurso.
- [x] **DI Container**: Configurar el contenedor de inyección de dependencias con `samber/do`.
- [x] **Main API**: Unificar todo en `cmd/api/main.go`.

## Batch 5: Integración y Polish (El Cierre)
- [x] **CORS & Logger**: Configurar middlewares globales en Echo.
- [x] **Manejo de Errores Global**: Asegurar que todos los errores devuelvan el formato JSON pactado para los Toasts del front.
- [x] **Variables de Entorno**: Refactorizar configuración (DB, Port, JWT) para usar `os.Getenv`.
- [x] **Pruebas de Integración**: Validar que el proyecto compila y los componentes están cableados (Tests unitarios en handlers).

## Batch 6: Integración Angular - Configuración (El Puente)
- [x] **Environments**: Configurar `src/environments/environment.ts` con la URL del backend de Go.
- [x] **HttpClient Module**: Asegurar que el proyecto use `provideHttpClient` o `HttpClientModule`.
- [x] **Auth Interceptor**: Crear un interceptor para adjuntar el JWT en cada petición a `/api/admin`.

## Batch 7: Integración Angular - Auth & Admin (La Puerta)
- [x] **Auth Service**: Implementar login contra el endpoint de Go y manejar el token con Signals.
- [x] **Auth Guard**: Proteger las rutas de edición en el front.
- [x] **Login UI**: Adaptar el componente de login para usar el nuevo servicio.

## Batch 8: Integración Angular - Datos (Los Músculos)
- [x] **Data Services**: Refactorizar los servicios de Proyectos, Experiencia y Educación para que llamen a la API de Go.
- [x] **Admin Dashboards**: Asegurar que los botones de "Crear/Editar/Borrar" funcionen con el nuevo backend.

## Batch 9: Integración Angular - Feedback (El Brillo)
- [x] **Toast Service**: Crear un servicio global de notificaciones usando Signals.
- [x] **Toast Component**: Implementar el componente visual para mostrar Success/Error/Info.
- [x] **Global Error Handling**: Integrar el manejador de errores del front con el Toast Service.
