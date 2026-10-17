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
)

func (s *userService) RefreshToken(ctx context.Context, request *dto.RefreshTokenRequest, userId string) (string, string, int, error) {
	refreshTokenExists, err := s.userRepository.GetUserByRefreshToken(ctx, request.RefreshToken, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("invalid or expired refresh token")
	}

	userId = refreshTokenExists.UserID

	userExist, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	if userExist == nil {
		return "", "", http.StatusUnauthorized, errors.New("user not found")
	}

	token, err := jwt.CreateToken(userId, userExist.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepository.DeleteRefreshToken(ctx, userId)

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()

	err = s.userRepository.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userId,
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
