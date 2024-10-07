package routers

import (
	"example.com/m/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

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
