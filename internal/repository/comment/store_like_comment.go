package comment

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func (r *commentRepository) StoreLikeComment(ctx context.Context, model *model.CommentLikeModel) error {
	commentLikeId := uuid.New().String()
	query := `INSERT INTO comment_likes (id, comment_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, commentLikeId, model.CommentID, model.UserID, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
