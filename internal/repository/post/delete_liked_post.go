package post

import "context"

func (r *postRepository) DeleteLikePost(ctx context.Context, postID string, userID string) error {
	query := `DELETE FROM post_likes WHERE post_id = ? and user_id = ?`

	_, err := r.db.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return err
	}
	return nil
}
