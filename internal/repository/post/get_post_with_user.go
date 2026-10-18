package post

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
)

func (r *postRepository) GetPostWithUser(ctx context.Context, postID string) (*model.PostWithUserModel, error) {
	query := `SELECT p.id,p.title,p.content,p.user_id,p.created_at,p.updated_at, u.username , COUNT(pl.id) as like_count
		FROM posts as p
		JOIN users as u ON p.user_id = u.id
		LEFT JOIN post_likes as pl ON p.id = pl.post_id
		WHERE p.id = ?
        GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username
		`

	row := r.db.QueryRowContext(ctx, query, postID)

	var post model.PostWithUserModel
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt, &post.Username, &post.LikeCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}
