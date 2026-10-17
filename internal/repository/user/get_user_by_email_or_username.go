package user

import (
	"TwitSocial/internal/model"
	"context"
	"database/sql"
	"errors"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error) {

	query := `
		SELECT * FROM users
		WHERE email = ? OR username = ?
	`

	row := r.db.QueryRowContext(ctx, query, email, username)
	var result model.UserModel
	err := row.Scan(&result.ID, &result.Email, &result.Password, &result.Username, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil

}
