package post

import (
	"TwitSocial/internal/middleware"
	"TwitSocial/internal/service/post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	postService post.PostService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, postService post.PostService) *Handler {
	return &Handler{
		api:         api,
		validate:    validate,
		postService: postService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/post")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/create", h.CreatePost)
	routeAuth.PUT("/update/:postId", h.UpdatePost)
	routeAuth.DELETE("/delete/:postId", h.DeletePost)
	routeAuth.POST("/action", h.LikeOrUnlikePost)

	routeWithoutAuth := h.api.Group("/post")
	routeWithoutAuth.GET("/detail/:postId", h.DetailPost)
}
