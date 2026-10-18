package comment

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (r *commentRepository) GetCommentByPostIDs(ctx context.Context, postIDs []string) ([]model.CommentWithLikeModel, error) {
	if len(postIDs) == 0 {
		return []model.CommentWithLikeModel{}, nil
	}

	placeholders := make([]string, len(postIDs))
	args := make([]interface{}, len(postIDs))
	for i, id := range postIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`SELECT c.id, c.post_id, c.user_id, u.username, c.content, c.created_at, c.updated_at, COUNT(cl.id) as like_count
                          FROM comments as c
                          JOIN users as u ON u.id = c.user_id
                          LEFT JOIN comment_likes as cl ON cl.comment_id = c.id
                          WHERE c.post_id IN (%s)
                          GROUP BY c.id, c.post_id, c.user_id, u.username, c.content, c.created_at, c.updated_at
                          ORDER BY like_count DESC`, strings.Join(placeholders, ", "))
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []model.CommentWithLikeModel{}, nil
		}
		return []model.CommentWithLikeModel{}, err
	}

	result := make([]model.CommentWithLikeModel, 0)
	for rows.Next() {
		var data model.CommentWithLikeModel
		err = rows.Scan(&data.ID, &data.PostID, &data.UserID, &data.Username, &data.Content, &data.CreatedAt, &data.UpdatedAt, &data.LikeCount)
		if err != nil {
			return []model.CommentWithLikeModel{}, err
		}
		result = append(result, model.CommentWithLikeModel{
			ID:        data.ID,
			PostID:    data.PostID,
			UserID:    data.UserID,
			Username:  data.Username,
			Content:   data.Content,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			LikeCount: data.LikeCount,
		})
	}
	return result, nil
}
