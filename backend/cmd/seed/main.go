package main

import (
	"fmt"
	"log"
	"os"
	"portfolio-backend/internal/adapters/auth"
	"portfolio-backend/pkg/db"
	"time"

	"github.com/jmoiron/sqlx"
)

func main() {
	// 1. Get Admin Credentials from Env
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	if adminEmail == "" || adminPass == "" {
		log.Fatal("ERROR: ADMIN_EMAIL and ADMIN_PASSWORD environment variables are required")
	}

	// 2. Setup DB
	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("ERROR: Could not connect to DB: %v", err)
	}
	defer database.Close()

	// 3. Hash Password
	authService := auth.NewAuthService("temp", 1*time.Hour) // Secret doesn't matter for hashing
	hash, err := authService.HashPassword(adminPass)
	if err != nil {
		log.Fatalf("ERROR: Could not hash password: %v", err)
	}

	// 4. Insert User
	err = createAdmin(database, adminEmail, hash)
	if err != nil {
		log.Fatalf("ERROR: Could not create admin: %v", err)
	}

	fmt.Printf("✅ Admin user created successfully: %s\n", adminEmail)
}

func createAdmin(db *sqlx.DB, email, hash string) error {
	query := `
		INSERT INTO users (email, password_hash, role)
		VALUES ($1, $2, 'admin')
		ON CONFLICT (email) DO UPDATE 
		SET password_hash = EXCLUDED.password_hash
	`
	_, err := db.Exec(query, email, hash)
	return err
}
