package domain

import "time"

type SocialLink struct {
	ID        string    `json:"id" db:"id"`
	Platform  string    `json:"platform" db:"platform"`
	URL       string    `json:"url" db:"url"`
	IconName  string    `json:"iconName" db:"icon_name"`
	IsActive  bool      `json:"isActive" db:"is_active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
