package post

import (
	"context"
	"errors"
	"net/http"
)

func (s *postService) DeletePost(ctx context.Context, postID string, userID string) (int, error) {
	postExists, err := s.postRepository.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if postExists == nil {
		return http.StatusNotFound, errors.New("post not found")
	}
	if postExists.UserID != userID {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	err = s.postRepository.SoftDeletePost(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}
