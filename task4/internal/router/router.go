package router

import (
	"task4/internal/handler"
	"task4/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestLogger(), gin.Recovery())
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	// 公开接口
	r.GET("/posts", handler.ListPosts)
	r.GET("/posts/:id", handler.GetPost)
	r.GET("/posts/:id/comments", handler.ListComments)
	// 鉴权接口
	auth := r.Group("")
	auth.Use(middleware.Auth())
	{
		auth.POST("/posts", handler.CreatePost)
		auth.PUT("/posts/:id", handler.UpdatePost)
		auth.DELETE("/posts/:id", handler.DeletePost)

		auth.POST("/comments", handler.CreateComment)
		auth.DELETE("/comments/:id", handler.DeleteComment)
	}
	return r
}
