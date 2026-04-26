# Diseño Técnico: Go Backend (Arquitectura Hexagonal)

## 1. Esquema de Base de Datos (DDL)
Usaremos SQL puro con `sqlx` y migraciones manuales.

```sql
-- Migration: 001_init.up.sql

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'admin',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    hero_title TEXT NOT NULL,
    hero_subtitle TEXT NOT NULL,
    about_me TEXT NOT NULL,
    avatar_url TEXT,
    resume_url TEXT,
    social_links JSONB NOT NULL DEFAULT '[]',
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT,
    github_url TEXT,
    live_url TEXT,
    category TEXT,
    tech_stack TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE experiences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company TEXT NOT NULL,
    position TEXT NOT NULL,
    period TEXT NOT NULL,
    description TEXT NOT NULL,
    tech_stack TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE education (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    institution TEXT NOT NULL,
    degree TEXT NOT NULL,
    period TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## 2. Estructura de Directorios (Hexagonal)

```text
backend/
├── cmd/
│   └── api/
│       └── main.go            # Wire everything (DI) and start server
├── internal/
│   ├── core/
│   │   ├── domain/            # Plain Go structs (Entities)
│   │   ├── ports/             # Interfaces (Repository, Service)
│   │   └── services/          # Business logic implementations
│   ├── adapters/
│   │   ├── handlers/          # HTTP Handlers (Echo)
│   │   ├── repositories/      # Postgres implementations (sqlx)
│   │   └── auth/              # JWT & Bcrypt logic
│   └── config/                # Environment config (Viper/Env)
├── pkg/
│   └── db/                    # DB connection pool (pgx)
├── migrations/                # SQL migration files
├── go.mod
└── Makefile
```

## 3. Inyección de Dependencias (samber/do)

Usaremos `samber/do` para desacoplar la creación de objetos. El `main.go` será el único lugar donde se registren las implementaciones concretas.

Ejemplo de registro:
```go
do.Provide(i, func(i *do.Injector) (ports.ProjectRepository, error) {
    db := do.MustInvoke[*sqlx.DB](i)
    return repositories.NewPostgresProjectRepository(db), nil
})
```

## 4. Stack Técnico Final
- **Router**: Echo v4.
- **DB Driver**: pgx/v5.
- **SQL Helper**: sqlx.
- **Auth**: JWT-Go.
- **Validation**: Validator v10.
- **Migrations**: golang-migrate.
