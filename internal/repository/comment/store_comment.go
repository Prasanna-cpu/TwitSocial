package comment

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (r *commentRepository) StoreComment(ctx context.Context, model *model.CommentModel) error {
	commentID := uuid.New().String()
	query := `INSERT INTO comments (id, post_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, commentID, model.PostID, model.UserID, model.Content, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}
