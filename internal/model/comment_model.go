package model

import "time"

type (
	CommentModel struct {
		ID        string
		PostID    string
		UserID    string
		Content   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	CommentLikeModel struct {
		ID        string
		CommentID string
		UserID    string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	CommentWithLikeModel struct {
		ID        string
		PostID    string
		UserID    string
		Content   string
		LikeCount int64
		Username  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
