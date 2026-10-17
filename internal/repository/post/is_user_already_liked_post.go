package post

import (
	"context"
	"database/sql"
	"errors"
)

func (r *postRepository) IsUserAlreadyLikedPost(ctx context.Context, postID string, userID string) (bool, error) {
	query := `SELECT id FROM post_likes WHERE post_id = ? AND user_id = ?`

	var id string
	err := r.db.QueryRowContext(ctx, query, postID, userID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
