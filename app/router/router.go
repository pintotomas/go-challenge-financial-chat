package router

import (
	"go-challenge-financial-chat/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup returns the app router.
func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping")
	})

	user := r.Group("/users")
	{
		user.POST("", controller.CreateUser)
		user.GET("/:userID", controller.GetUser)
	}

	return r
}
