package main

import (
	"TwitSocial/internal/config"
	commentHandler "TwitSocial/internal/handler/comment"
	postHandler "TwitSocial/internal/handler/post"
	userHandler "TwitSocial/internal/handler/user"
	commentRepo "TwitSocial/internal/repository/comment"
	postRepo "TwitSocial/internal/repository/post"
	userRepo "TwitSocial/internal/repository/user"
	commentService "TwitSocial/internal/service/comment"
	postService "TwitSocial/internal/service/post"
	userService "TwitSocial/internal/service/user"
	"TwitSocial/pkg/internalsql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	db, err := internalsql.ConnectMySQL(loadConfig)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}

	userRepo := userRepo.NewRepository(db)
	postRepo := postRepo.NewPostRepository(db)
	commentRepo := commentRepo.NewCommentRepository(db)
	userService := userService.NewService(loadConfig, userRepo)
	postService := postService.NewPostService(loadConfig, postRepo, commentRepo)
	commentService := commentService.NewCommentService(loadConfig, commentRepo, postRepo)
	validate := validator.New()
	userHandler := userHandler.NewHandler(r, validate, userService)
	postHandler := postHandler.NewHandler(r, validate, postService)
	commentHandler := commentHandler.NewHandler(r, validate, commentService)
	userHandler.RouteList(loadConfig.SecretJwt)
	postHandler.RouteList(loadConfig.SecretJwt)
	commentHandler.RouteList(loadConfig.SecretJwt)

	server := fmt.Sprintf(":%s", loadConfig.Port)
	err = r.Run(server)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
