package middleware

import (
	"TwitSocial/pkg/jwt"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			context.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized, missing token"))
			return
		}
		if strings.HasPrefix(header, "Bearer ") {
			header = strings.TrimPrefix(header, "Bearer ")
		}
		//log.Println("Token after stripping Bearer:", header)
		userID, username, err := jwt.ValidateToken(header, secretKey, true)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized, invalid token"))
			return
		}
		//log.Println("Setting userID:", userID, "username:", username)
		context.Set("userID", userID)
		context.Set("username", username)
		context.Next()
	}
}

func AuthRefreshTokenMiddleware(secretKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			context.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized, missing token"))
			return
		}
		if strings.HasPrefix(header, "Bearer ") {
			header = strings.TrimPrefix(header, "Bearer ")
		}
		userID, username, err := jwt.ValidateToken(header, secretKey, false)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized, invalid token"))
			return
		}
		context.Set("userID", userID)
		context.Set("username", username)
		context.Next()
	}
}
