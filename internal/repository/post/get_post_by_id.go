package post

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
)

func (r *postRepository) GetPostByID(ctx context.Context, postID string) (*model.PostModel, error) {
	query := `SELECT id,title,content,user_id,created_at,updated_at
		FROM posts
		WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, postID)

	var post model.PostModel
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}
