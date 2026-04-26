package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/ports"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/do/v2"
)

// SetupRouter configures Echo routes and middlewares.
func SetupRouter(e *echo.Echo, injector do.Injector) {
	// Global Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	// Serve Static Files (Uploads)
	e.Static("/uploads", "uploads")

	// Invoke Services from DI
	authAppService := do.MustInvoke[ports.AuthAppService](injector)
	authAdapter := do.MustInvoke[ports.AuthService](injector)
	projectService := do.MustInvoke[ports.ProjectService](injector)
	experienceService := do.MustInvoke[ports.ExperienceService](injector)
	educationService := do.MustInvoke[ports.EducationService](injector)
	heroService := do.MustInvoke[ports.HeroService](injector)
	skillService := do.MustInvoke[ports.SkillService](injector)
	socialService := do.MustInvoke[ports.SocialService](injector)

	// Handlers
	authH := NewAuthHandler(authAppService)
	projectH := NewProjectHandler(projectService)
	experienceH := NewExperienceHandler(experienceService)
	educationH := NewEducationHandler(educationService)
	heroH := NewHeroHandler(heroService)
	skillH := NewSkillHandler(skillService)
	socialH := NewSocialHandler(socialService)
	mediaH := NewMediaHandler("uploads")

	// Public Routes
	e.POST("/api/auth/login", authH.Login)
	e.GET("/api/projects", projectH.GetAll)
	e.GET("/api/projects/:id", projectH.GetByID)
	e.GET("/api/experience", experienceH.GetAll)
	e.GET("/api/education", educationH.GetAll)
	e.GET("/api/hero", heroH.Get)
	e.GET("/api/skills", skillH.GetAll)
	e.GET("/api/social", socialH.GetAll)

	// Protected Routes (Admin only)
	admin := e.Group("/api/admin")
	admin.Use(AuthMiddleware(authAdapter))

	// Projects
	admin.POST("/projects", projectH.Create)
	admin.PUT("/projects/:id", projectH.Update)
	admin.DELETE("/projects/:id", projectH.Delete)

	// Experience
	admin.POST("/experience", experienceH.Create)
	admin.PUT("/experience/:id", experienceH.Update)
	admin.DELETE("/experience/:id", experienceH.Delete)

	// Education
	admin.POST("/education", educationH.Create)
	admin.PUT("/education/:id", educationH.Update)
	admin.DELETE("/education/:id", educationH.Delete)

	// Hero
	admin.PUT("/hero", heroH.Update)

	// Skills
	admin.POST("/skills", skillH.Create)
	admin.PUT("/skills/:id", skillH.Update)
	admin.DELETE("/skills/:id", skillH.Delete)

	// Social Links
	admin.PUT("/social/:id", socialH.Update)

	// Media
	admin.POST("/upload", mediaH.Upload)
}
