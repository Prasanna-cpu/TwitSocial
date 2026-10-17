package post

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func (r *postRepository) StorePost(ctx context.Context, model *model.PostModel) (string, error) {
	postId := uuid.New().String()
	query := "INSERT INTO posts (id, user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, postId, model.UserID, model.Title, model.Content, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("no rows")
		}
		return "", err
	}

	return postId, nil
}
