# Exploración: Backend en Go con Arquitectura Hexagonal

## Objetivo
Diseñar un backend robusto, escalable y ultra optimizado para el Portfolio, utilizando Go, Postgres y Arquitectura Hexagonal, con inyección de dependencias mediante `samber/do`.

## Investigaciones

### Samber/do en Arquitectura Hexagonal
`samber/do` es una librería de DI basada en genéricos que encaja perfecto con Go moderno. 
- **Ventaja**: Evita el acoplamiento manual y facilita el testing (pudiendo mockear interfaces fácilmente).
- **Implementación**: Registraremos los adaptadores (DB, Auth) y servicios en el `main.go`. Los servicios recibirán sus dependencias a través de interfaces (ports).

### Estructura del Proyecto
Adoptaremos una estructura Hexagonal (Ports & Adapters) para separar la lógica de negocio de la infraestructura:
- `internal/core/domain`: Entidades puras (Project, Experience, User). Sin dependencias externas.
- `internal/core/ports`: Interfaces que definen cómo se comunica el core con el mundo exterior.
- `internal/core/services`: Implementación de la lógica de negocio (Casos de uso).
- `internal/adapters`: Implementaciones concretas (Postgres, JWT, HTTP Handlers).

## Diseño Propuesto

### Estructura de Directorios
```
backend/
├── cmd/
│   └── api/
│       └── main.go            # DI setup & Server start
├── internal/
│   ├── core/
│   │   ├── domain/            # Entities (Project, Experience, User)
│   │   ├── ports/             # Interfaces (Repository, Service)
│   │   └── services/          # Business logic
│   ├── adapters/
│   │   ├── handlers/          # REST Handlers
│   │   ├── repositories/      # Postgres implementation
│   │   └── auth/              # JWT logic
│   └── config/                # Configuration
├── pkg/
│   ├── logger/                # Logging helper
│   └── database/              # DB connection
├── migrations/                # SQL scripts
├── go.mod
└── Makefile
```

### Esquema de Base de Datos (Postgres)
- **Users**: Para la gestión administrativa (Auth).
- **Projects**: Títulos, descripciones, links, tags (usando arrays de Postgres para performance).
- **Experiences**: Historial laboral.

### Autenticación
- JWT (JSON Web Tokens).
- Middleware de protección para rutas sensibles (POST, PUT, DELETE).
- El admin será el único con permisos de escritura.

## Próximos Pasos
1. Definir los contratos (Ports) iniciales.
2. Configurar el contenedor de DI (`samber/do`).
3. Implementar el adaptador de Postgres.
