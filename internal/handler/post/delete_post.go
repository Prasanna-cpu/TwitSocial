package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeletePost(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		postID = c.Param("postId")
		userID = c.GetString("userID")
	)

	statusCode, err := h.postService.DeletePost(ctx, postID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "Post deleted successfully",
	})

}
