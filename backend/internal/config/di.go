package config

import (
	"portfolio-backend/internal/adapters/auth"
	"portfolio-backend/internal/adapters/repositories"
	"portfolio-backend/internal/core/ports"
	"portfolio-backend/internal/core/services"
	"portfolio-backend/pkg/db"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do/v2"
)

// SetupDI initializes the dependency injection container.
func SetupDI() do.Injector {
	injector := do.New()

	// 1. Infrastructure
	do.Provide(injector, func(i do.Injector) (*sqlx.DB, error) {
		return db.NewPostgresDB()
	})

	// 2. Adapters (External)
	do.Provide(injector, func(i do.Injector) (ports.AuthService, error) {
		secret := GetEnv("JWT_SECRET", "super-secret-key-change-me")
		return auth.NewAuthService(secret, 24*time.Hour), nil
	})

	// 3. Repositories
	do.Provide(injector, func(i do.Injector) (ports.UserRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewPostgresUserRepository(dbInstance), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.ProjectRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewPostgresProjectRepository(dbInstance), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.ExperienceRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewExperienceRepository(dbInstance), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.EducationRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewEducationRepository(dbInstance), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.HeroRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewHeroRepository(dbInstance), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.SkillRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewSkillRepository(dbInstance), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.SocialRepository, error) {
		dbInstance := do.MustInvoke[*sqlx.DB](i)
		return repositories.NewSocialRepository(dbInstance), nil
	})

	// 4. Application Services
	do.Provide(injector, func(i do.Injector) (ports.AuthAppService, error) {
		userRepo := do.MustInvoke[ports.UserRepository](i)
		authAdapter := do.MustInvoke[ports.AuthService](i)
		return services.NewAuthAppService(userRepo, authAdapter), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.ProjectService, error) {
		repo := do.MustInvoke[ports.ProjectRepository](i)
		return services.NewProjectService(repo), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.ExperienceService, error) {
		repo := do.MustInvoke[ports.ExperienceRepository](i)
		return services.NewExperienceService(repo), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.EducationService, error) {
		repo := do.MustInvoke[ports.EducationRepository](i)
		return services.NewEducationService(repo), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.HeroService, error) {
		repo := do.MustInvoke[ports.HeroRepository](i)
		return services.NewHeroService(repo), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.SkillService, error) {
		repo := do.MustInvoke[ports.SkillRepository](i)
		return services.NewSkillService(repo), nil
	})
	do.Provide(injector, func(i do.Injector) (ports.SocialService, error) {
		repo := do.MustInvoke[ports.SocialRepository](i)
		return services.NewSocialService(repo), nil
	})

	return injector
}
