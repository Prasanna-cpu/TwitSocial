package user

import (
	"TwitSocial/internal/model"
	"context"

	"github.com/google/uuid"
)

func (r *userRepository) CreateUser(ctx context.Context, model *model.UserModel) (string, error) {

	// Generate UUID in application code
	userID := uuid.New().String()

	query := `
		INSERT INTO users (id, email, username, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, userID, model.Email, model.Username, model.Password, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return "", err
	}

	return userID, nil
}
