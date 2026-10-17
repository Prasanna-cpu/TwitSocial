package user

import (
	"TwitSocial/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request dto.RefreshTokenRequest
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

	token, refreshToken, statusCode, err := h.userService.RefreshToken(ctx, &request, "")
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})

}
