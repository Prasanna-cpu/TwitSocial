package comment

import (
	"TwitSocial/internal/config"
	"TwitSocial/internal/dto"
	"TwitSocial/internal/repository/comment"
	"TwitSocial/internal/repository/post"
	"context"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID string) (int, error)
	LikeOrUnlikeComment(ctx context.Context, commentID string, userID string) (int, error)
}

type commentService struct {
	cfg               *config.Config
	commentRepository comment.CommentRepository
	postRepository    post.PostRepository
}

func NewCommentService(cfg *config.Config, commentRepository comment.CommentRepository, postRepository post.PostRepository) CommentService {
	return &commentService{
		cfg:               cfg,
		commentRepository: commentRepository,
		postRepository:    postRepository,
	}
}
