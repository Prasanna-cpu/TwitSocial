package post

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) DetailPost(c *gin.Context) {
	ctx := c.Request.Context()
	postID := c.Param("postId")

	result, statuscode, err := h.postService.DetailPost(ctx, postID)
	if err != nil {
		c.JSON(statuscode, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(statuscode, result)
}
