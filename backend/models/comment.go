package models

import (
	"time"
)

type Comment struct {
	CommentID       int       `json:"comment_id" db:"comment_id"`
	VideoID         int       `json:"video_id" db:"video_id"`
	UserID          int       `json:"user_id" db:"user_id"`
	CommentText     string    `json:"comment_text" db:"comment_text"`
	CommentDate     time.Time `json:"comment_date" db:"comment_date"`
	ParentCommentID *int      `json:"parent_comment_id,omitempty" db:"parent_comment_id"` // Ponteiro para int para permitir nulo no DB
}
