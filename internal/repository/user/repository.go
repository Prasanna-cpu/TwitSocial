package user

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"time"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (string, error)
	GetRefreshToken(ctx context.Context, id string, now time.Time) (*model.RefreshTokenModel, error)
	StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error
	GetUserById(ctx context.Context, id string) (*model.UserModel, error)
	DeleteRefreshToken(ctx context.Context, userId string) error
	GetUserByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*model.RefreshTokenModel, error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
