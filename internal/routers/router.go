package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong) // "/v1/2024/ping"
	}

	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", Pong) // "/v2/2024/ping"
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "NDD")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"users":   []string{"str1", "str2"},
	})
}
