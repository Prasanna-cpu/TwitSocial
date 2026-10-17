package comment

import (
	"context"
	"database/sql"
	"errors"
)

func (r *commentRepository) IsUserAlreadyLikedComment(ctx context.Context, commentID string, userID string) (bool, error) {
	query := `SELECT id FROM comment_likes WHERE comment_id = ? AND user_id = ?`

	row := r.db.QueryRowContext(ctx, query, commentID, userID)
	var id string
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
