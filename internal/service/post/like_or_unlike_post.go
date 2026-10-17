package post

import (
	"TwitSocial/internal/model"
	"context"
	"errors"
	"net/http"
	"time"
)

func (s *postService) LikeOrUnlikePost(ctx context.Context, postID string, userID string) (int, error) {
	postExist, err := s.postRepository.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if postExist == nil {
		return http.StatusNotFound, errors.New("post not found")
	}

	isUserAlreadyLikedPost, err := s.postRepository.IsUserAlreadyLikedPost(ctx, postID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if isUserAlreadyLikedPost {
		err := s.postRepository.DeleteLikePost(ctx, postID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		now := time.Now()
		err := s.postRepository.StoreLikePost(ctx, &model.PostLikeModel{
			PostID:    postID,
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
