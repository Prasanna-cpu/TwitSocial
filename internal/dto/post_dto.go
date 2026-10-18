package dto

import "time"

type (
	CreateOrUpdatePostRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	CreateOrUpdatePostResponse struct {
		ID string `json:"id"`
	}
)

type (
	LikeOrUnlikePostRequest struct {
		PostID string `json:"post_id" validate:"required"`
	}
)

type (
	Comment struct {
		ID        string    `json:"id"`
		Username  string    `json:"username"`
		Content   string    `json:"content"`
		LikeCount int64     `json:"like_count"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	DetailPostResponse struct {
		ID        string    `json:"id"`
		Username  string    `json:"username"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		LikeCount int64     `json:"like_count"`
		Comments  []Comment `json:"comments"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
