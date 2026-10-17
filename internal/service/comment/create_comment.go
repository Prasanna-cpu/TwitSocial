package comment

import (
	"TwitSocial/internal/dto"
	"TwitSocial/internal/model"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s *commentService) CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID string) (int, error) {
	postExists, err := s.postRepository.GetPostByID(ctx, req.PostID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if postExists == nil {
		return http.StatusNotFound, fmt.Errorf("post not found")
	}
	now := time.Now()

	err = s.commentRepository.StoreComment(ctx, &model.CommentModel{
		PostID:    req.PostID,
		UserID:    userID,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}
