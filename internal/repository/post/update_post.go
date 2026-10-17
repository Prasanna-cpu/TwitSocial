package post

import (
	"TwitSocial/internal/model"
	"context"
	"errors"
)

func (r *postRepository) UpdatePost(ctx context.Context, model *model.PostModel, postID string) error {
	query := `UPDATE posts SET title = ?, content = ?, updated_at = ? WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, model.Title, model.Content, model.UpdatedAt, postID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
