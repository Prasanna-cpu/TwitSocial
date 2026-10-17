package user

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
	"time"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, id string, now time.Time) (*model.RefreshTokenModel, error) {
	query := `SELECT id , user_id , refresh_token , expired_at , created_at , updated_at
	FROM refresh_tokens
	WHERE user_id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	var result model.RefreshTokenModel
	err := row.Scan(&result.ID, &result.UserID, &result.RefreshToken, &result.ExpiredAt, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil

}
