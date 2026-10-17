package post

import (
	"TwitSocial/internal/dto"
	"TwitSocial/internal/model"
	"context"
	"errors"
	"net/http"
	"time"
)

func (s *postService) UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID string, postID string) (string, int, error) {
	postExists, err := s.postRepository.GetPostByID(ctx, postID)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	if postExists == nil {
		return "", http.StatusNotFound, errors.New("post not found")
	}
	if postExists.UserID != userID {
		return "", http.StatusUnauthorized, errors.New("unauthorized")
	}

	err = s.postRepository.UpdatePost(ctx, &model.PostModel{
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}, postID)

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return postID, http.StatusOK, nil

}
