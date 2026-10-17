package model

import "time"

type (
	UserModel struct {
		ID        string
		Email     string
		Username  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefreshTokenModel struct {
		ID           string
		UserID       string
		RefreshToken string
		ExpiredAt    time.Time
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
