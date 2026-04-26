# Fullstack Dynamic Portfolio & CMS

![Angular](https://img.shields.io/badge/Angular-v21-dd0031?style=for-the-badge&logo=angular)
![Go](https://img.shields.io/badge/Go-1.26-00add8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?style=for-the-badge&logo=postgresql)

A professional, high-performance portfolio application featuring a custom-built Content Management System (CMS). Built with **Hexagonal Architecture** in the backend and **Angular v21 (Signals & Native Control Flow)** in the frontend.

## 🚀 Key Features

- **Dynamic CMS**: Manage Hero section, Skills, Social Links, Professional Experience, and Education without touching the code.
- **Media Management**: Upload and manage profile images, project screenshots, and CV files directly from the UI.
- **Modern UI/UX**: Premium design with glassmorphism, smooth gradients, and fully responsive layouts.
- **Secure Authentication**: JWT-based admin authentication to protect CMS actions.
- **Advanced State Management**: Powered by Angular Signals for reactive and efficient UI updates.

## 🛠️ Tech Stack

### Backend
- **Language**: Go (Golang)
- **Framework**: Echo (High performance, minimalist web framework)
- **Architecture**: Hexagonal (Ports & Adapters) for maximum testability and decoupling.
- **Database**: PostgreSQL with `sqlx` and `pgx` driver.
- **DI**: Samber/DO for robust Dependency Injection.

### Frontend
- **Framework**: Angular v21 (Standalone Components)
- **Logic**: Signals for state management and reactive control flow (`@if`, `@for`).
- **Styling**: Vanilla SCSS with a custom design system.
- **Features**: Reactive Forms, NgOptimizedImage, and Modular Modals.

## 📦 Setup & Installation

### Prerequisites
- Go 1.23+
- Node.js 20+
- PostgreSQL 15+ (or Docker)

### 1. Database Setup
You can use the provided `docker-compose.yaml` to spin up the database and pgAdmin:
```bash
cd backend
docker-compose up -d
```

### 2. Backend Configuration
The backend uses environment variables. You can set them in your terminal or a `.env` file (not committed).

Run migrations (if you have a migration tool) or ensure the schema is loaded from `backend/migrations/`.

### 3. Create Admin User (CRITICAL)
Since the repository is clean of hardcoded credentials, you must create your first admin user manually:
```bash
cd backend
ADMIN_EMAIL=your@email.com ADMIN_PASSWORD=your_password go run cmd/seed/main.go
```

### 4. Run the Project

**Backend:**
```bash
cd backend
go run cmd/api/main.go
```

**Frontend:**
```bash
cd frontend
npm install
npm start
```

## 📂 Project Structure
```text
.
├── backend/            # Go Clean Architecture Backend
│   ├── cmd/            # Entry points (API & Seeder)
│   ├── internal/       # Core business logic & Adapters
│   └── migrations/     # Database schema versions
├── frontend/           # Angular v21 Application
│   ├── src/app/        # Components, Services, and Signals
│   └── src/styles/     # Global design system
└── openspec/           # SDD (Spec-Driven Development) Artifacts
```

## 📄 License
This project is open-source and available under the MIT License.

---
Developed with ❤️ by [Salvucci Facundo](https://github.com/SalvucciFacundo)
