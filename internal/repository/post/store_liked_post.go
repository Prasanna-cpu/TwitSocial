package post

import (
	"TwitSocial/internal/model"
	"context"

	"github.com/google/uuid"
)

func (r *postRepository) StoreLikePost(ctx context.Context, model *model.PostLikeModel) error {
	postlikeId := uuid.New().String()

	query := `INSERT INTO post_likes (id, post_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, postlikeId, model.PostID, model.UserID, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
