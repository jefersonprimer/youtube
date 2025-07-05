package models

import (
	"time"
)

type Video struct {
	VideoID         int       `json:"video_id" db:"video_id"`
	UserID          int       `json:"user_id" db:"user_id"`
	Title           string    `json:"title" db:"title"`
	Description     *string   `json:"description,omitempty" db:"description"`
	FileUrl         string    `json:"file_url" db:"file_url"`
	ThumbnailUrl    string    `json:"thumbnail_url" db:"thumbnail_url"`
	UploadDate      time.Time `json:"upload_date" db:"upload_date"`
	Views           int       `json:"views" db:"views"`
	DurationSeconds *int      `json:"duration_seconds,omitempty" db:"duration_seconds"` // Ponteiro para int para permitir nulo no DB
	Visibility      string    `json:"visibility" db:"visibility"`                       // 'public', 'unlisted', 'private'
}
