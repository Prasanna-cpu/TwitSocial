package post

import (
	"TwitSocial/internal/config"
	"TwitSocial/internal/dto"
	"TwitSocial/internal/repository/post"
	"context"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID string) (string, int, error)
	UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID string, postID string) (string, int, error)
	DeletePost(ctx context.Context, postID string, userID string) (int, error)
	LikeOrUnlikePost(ctx context.Context, postID string, userID string) (int, error)
}

type postService struct {
	cfg            *config.Config
	postRepository post.PostRepository
}

func NewPostService(cfg *config.Config, postRepository post.PostRepository) PostService {
	return &postService{
		cfg:            cfg,
		postRepository: postRepository,
	}
}
