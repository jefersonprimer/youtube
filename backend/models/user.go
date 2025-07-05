package models

import (
	"time"
)

type User struct {
	UserID            int       `json:"user_id" db:"user_id"`
	Username          string    `json:"username" db:"username"`
	Email             string    `json:"email" db:"email"`
	PasswordHash      string    `json:"password_hash" db:"password_hash"`
	JoinDate          time.Time `json:"join_date" db:"join_date"`
	ProfilePictureURL *string   `json:"profile_picture_url,omitempty" db:"profile_picture_url"` // Ponteiro para string para permitir nulo no DB
	ChannelName       *string   `json:"channel_name,omitempty" db:"channel_name"`
	Description       *string   `json:"description,omitempty" db:"description"`
}
