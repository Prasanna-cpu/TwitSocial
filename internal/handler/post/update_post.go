package post

import (
	"TwitSocial/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdatePost(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request dto.CreateOrUpdatePostRequest
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
	postID := c.Param("postId")
	insertedId, statusCode, err := h.postService.UpdatePost(ctx, &request, userID, postID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "Post updated successfully",
		"id":      insertedId,
	})
}
