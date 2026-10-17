package user

import (
	"TwitSocial/internal/dto"
	"TwitSocial/internal/model"
	"TwitSocial/pkg/jwt"
	"TwitSocial/pkg/refreshtoken"
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, request *dto.LoginRequest) (string, string, int, error) {
	userExists, err := s.userRepository.GetUserByEmailOrUsername(ctx, request.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(request.Password))
	if err != nil {
		return "", "", http.StatusUnauthorized, err
	}

	token, err := jwt.CreateToken(userExists.ID, userExists.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()
	refreshTokenExist, err := s.userRepository.GetRefreshToken(ctx, userExists.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExist != nil {
		return token, refreshTokenExist.RefreshToken, http.StatusOK, nil
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepository.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userExists.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
	})

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refreshToken, http.StatusOK, nil

}
