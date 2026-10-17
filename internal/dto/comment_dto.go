package dto

type (
	StoreCommentRequest struct {
		PostID  string `json:"post_id" validate:"required"`
		Content string `json:"content" validate:"required"`
	}
)

type (
	LikeOrUnlikeCommentRequest struct {
		CommentID string `json:"comment_id" validate:"required"`
	}
)
