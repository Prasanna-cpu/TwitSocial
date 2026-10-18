package model

import "time"

type (
	PostModel struct {
		ID        string
		UserID    string
		Title     string
		Content   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}

	PostLikeModel struct {
		ID        string
		PostID    string
		UserID    string
		CreatedAt time.Time
		UpdatedAt time.Time
		//DeletedAt time.Time
	}

	PostWithUserModel struct {
		ID        string
		UserID    string
		Username  string
		Title     string
		Content   string
		LikeCount int64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
