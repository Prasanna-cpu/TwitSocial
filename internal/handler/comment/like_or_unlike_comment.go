package comment

import (
	"TwitSocial/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LikeOrUnlikeComment(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request dto.LikeOrUnlikeCommentRequest
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetString("userID")
	//postID := c.Param("postId")
	statusCode, err := h.commentService.LikeOrUnlikeComment(ctx, request.CommentID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "Comment liked successfully",
	})
}
