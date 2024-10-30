package initialize

import (
	"fmt"

	"example.com/m/internal/controller"
	"example.com/m/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// middleware
func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

// middleware
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

// middleware
func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", controller.NewPongController().Pong)        // "/v1/2024/ping"
		v1.GET("/user", controller.NewUserController().GetUserById) // "/v1/2024/ping"
	}

	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", controller.NewPongController().Pong) // "/v2/2024/ping"
	}

	return r
}
