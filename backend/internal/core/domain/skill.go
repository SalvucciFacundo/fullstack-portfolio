package domain

import "time"

type Skill struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	IconClass    string    `json:"iconClass" db:"icon_class"`
	Category     string    `json:"category" db:"category"`
	DisplayOrder int       `json:"displayOrder" db:"display_order"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}
