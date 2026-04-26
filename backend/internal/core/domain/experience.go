package domain

import "time"

// Experience represents a work experience entry.
type Experience struct {
	ID          string     `json:"id" db:"id"`
	Company     string     `json:"company" db:"company"`
	Role        string     `json:"role" db:"role"`
	Description string     `json:"description" db:"description"`
	StartDate   time.Time  `json:"startDate" db:"start_date"`
	EndDate     *time.Time `json:"endDate" db:"end_date"`
	IsCurrent   bool       `json:"isCurrent" db:"is_current"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
}
