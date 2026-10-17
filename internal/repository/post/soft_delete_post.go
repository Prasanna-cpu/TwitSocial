package post

import (
	"context"
	"errors"
)

func (r *postRepository) SoftDeletePost(ctx context.Context, postID string) error {
	query := `DELETE from posts where id = ?`
	result, err := r.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
