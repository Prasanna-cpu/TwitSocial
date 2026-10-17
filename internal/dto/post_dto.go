package dto

type (
	CreateOrUpdatePostRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	CreateOrUpdatePostResponse struct {
		ID string `json:"id"`
	}
)

type (
	LikeOrUnlikePostRequest struct {
		PostID string `json:"post_id" validate:"required"`
	}
)
