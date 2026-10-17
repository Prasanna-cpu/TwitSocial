package user

import "context"

func (r *userRepository) DeleteRefreshToken(ctx context.Context, userId string) error {
	query := `DELETE FROM refresh_tokens where user_id = ?`
	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}
	return nil
}
