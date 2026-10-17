package comment

import (
	"TwitSocial/internal/model"
	"context"
	"errors"
	"net/http"
	"time"
)

func (s *commentService) LikeOrUnlikeComment(ctx context.Context, commentID string, userID string) (int, error) {
	commentExists, err := s.commentRepository.DetailCommentByID(ctx, commentID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if commentExists == nil {
		return http.StatusNotFound, errors.New("comment not found")
	}

	isLiked, err := s.commentRepository.IsUserAlreadyLikedComment(ctx, commentID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if isLiked {
		err = s.commentRepository.DeleteLikeComment(ctx, commentID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		now := time.Now()
		err = s.commentRepository.StoreLikeComment(ctx, &model.CommentLikeModel{
			CommentID: commentID,
			UserID:    userID,
			CreatedAt: now,
			UpdatedAt: now,
		})
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
