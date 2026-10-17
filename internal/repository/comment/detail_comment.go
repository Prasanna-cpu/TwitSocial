package comment

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
)

func (r *commentRepository) DetailCommentByID(ctx context.Context, commentID string) (*model.CommentModel, error) {
	query := `SELECT id, post_id, user_id, content, created_at, updated_at FROM comments WHERE id = ?`
	rows := r.db.QueryRowContext(ctx, query, commentID)
	var result model.CommentModel
	err := rows.Scan(&result.ID, &result.PostID, &result.UserID, &result.Content, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
