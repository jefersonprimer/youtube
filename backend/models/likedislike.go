package models

import (
	"time"
)

type LikeDislike struct {
	LikeID       int       `json:"like_id" db:"like_id"`
	UserID       int       `json:"user_id" db:"user_id"`
	VideoID      *int      `json:"video_id,omitempty" db:"video_id"`     // Ponteiro para int, pois pode ser nulo (se for em comentário)
	CommentID    *int      `json:"comment_id,omitempty" db:"comment_id"` // Ponteiro para int, pois pode ser nulo (se for em vídeo)
	Type         string    `json:"type" db:"type"`                       // 'like' ou 'dislike'
	ReactionDate time.Time `json:"reaction_date" db:"reaction_date"`
}
