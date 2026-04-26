package domain

import "time"

// Education represents an academic qualification.
type Education struct {
	ID          string     `json:"id" db:"id"`
	Institution string     `json:"institution" db:"institution"`
	Degree      string     `json:"degree" db:"degree"`
	StartDate   time.Time  `json:"startDate" db:"start_date"`
	EndDate     *time.Time `json:"endDate" db:"end_date"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
}
