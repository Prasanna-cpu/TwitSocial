package post

import (
	"TwitSocial/internal/dto"
	"context"
	"errors"
	"net/http"
)

func (s *postService) DetailPost(ctx context.Context, postID string) (*dto.DetailPostResponse, int, error) {
	post, err := s.postRepository.GetPostWithUser(ctx, postID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if post == nil {
		return nil, http.StatusNotFound, errors.New("post not found")
	}
	postIDs := []string{post.ID}
	comments, err := s.commentRepository.GetCommentByPostIDs(ctx, postIDs)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	commentsMap := make([]dto.Comment, 0)
	for _, comment := range comments {
		commentsMap = append(commentsMap, dto.Comment{
			ID:        comment.ID,
			Username:  comment.Username,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			LikeCount: comment.LikeCount,
		})
	}

	return &dto.DetailPostResponse{
		ID:        post.ID,
		Username:  post.Username,
		Title:     post.Title,
		LikeCount: post.LikeCount,
		Comments:  commentsMap,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, http.StatusOK, nil

}
