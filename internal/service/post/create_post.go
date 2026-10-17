package post

import (
	"TwitSocial/internal/dto"
	"TwitSocial/internal/model"
	"context"
	"net/http"
	"time"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID string) (string, int, error) {
	now := time.Now()
	insertedId, err := s.postRepository.StorePost(ctx, &model.PostModel{
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return insertedId, http.StatusCreated, nil

}
