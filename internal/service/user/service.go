package user

import (
	"TwitSocial/internal/config"
	"TwitSocial/internal/dto"
	"TwitSocial/internal/repository/user"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request *dto.RegisterRequest) (string, int, error)
	Login(ctx context.Context, request *dto.LoginRequest) (string, string, int, error)
	RefreshToken(ctx context.Context, request *dto.RefreshTokenRequest, userId string) (string, string, int, error)
}

type userService struct {
	cfg            *config.Config
	userRepository user.UserRepository
}

func NewService(cfg *config.Config, userRepository user.UserRepository) UserService {
	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
	}
}
