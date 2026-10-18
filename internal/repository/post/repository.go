package post

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
)

type PostRepository interface {
	StorePost(ctx context.Context, model *model.PostModel) (string, error)
	GetPostByID(ctx context.Context, postID string) (*model.PostModel, error)
	UpdatePost(ctx context.Context, model *model.PostModel, postID string) error
	SoftDeletePost(ctx context.Context, postID string) error
	IsUserAlreadyLikedPost(ctx context.Context, postID string, userID string) (bool, error)
	DeleteLikePost(ctx context.Context, postID string, userID string) error
	StoreLikePost(ctx context.Context, model *model.PostLikeModel) error
	GetPostWithUser(ctx context.Context, postID string) (*model.PostWithUserModel, error)
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
