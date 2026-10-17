package comment

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
)

type CommentRepository interface {
	StoreComment(ctx context.Context, model *model.CommentModel) error
	DetailCommentByID(ctx context.Context, commentID string) (*model.CommentModel, error)
	IsUserAlreadyLikedComment(ctx context.Context, commentID string, userID string) (bool, error)
	DeleteLikeComment(ctx context.Context, commentID string, userID string) error
	StoreLikeComment(ctx context.Context, model *model.CommentLikeModel) error
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}
