package domain

import "time"

type HeroSection struct {
	ID           string    `json:"id" db:"id"`
	Headline     string    `json:"headline" db:"headline"`
	Subheadline  string    `json:"subheadline" db:"subheadline"`
	Biography    string    `json:"biography" db:"biography"`
	ProfileImage string    `json:"profileImage" db:"profile_image"`
	ResumeURL    string    `json:"resumeUrl" db:"resume_url"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}
