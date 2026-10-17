package user

import (
	"TwitSocial/internal/model"
	"context"

	"github.com/google/uuid"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error {

	newRefreshTokenID := uuid.New().String()

	query := `INSERT INTO refresh_tokens (id, user_id, refresh_token, expired_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, newRefreshTokenID, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
