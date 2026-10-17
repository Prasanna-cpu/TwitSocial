package user

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
)

func (r *userRepository) GetUserById(ctx context.Context, id string) (*model.UserModel, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	user := &model.UserModel{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
