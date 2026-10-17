package post

import (
	"TwitSocial/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LikeOrUnlikePost(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request dto.LikeOrUnlikePostRequest
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
	statusCode, err := h.postService.LikeOrUnlikePost(ctx, request.PostID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "Post liked successfully",
	})
}
