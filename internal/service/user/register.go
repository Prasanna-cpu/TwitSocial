package user

import (
	"TwitSocial/internal/dto"
	"TwitSocial/internal/model"
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, request *dto.RegisterRequest) (string, int, error) {
	userExist, err := s.userRepository.GetUserByEmailOrUsername(ctx, request.Email, request.Username)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	if userExist != nil {
		return "", http.StatusConflict, errors.New("user already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	now := time.Now()
	userModel := &model.UserModel{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(password),
		CreatedAt: now,
		UpdatedAt: now,
	}

	userId, err := s.userRepository.CreateUser(ctx, userModel)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return userId, http.StatusCreated, nil

}
