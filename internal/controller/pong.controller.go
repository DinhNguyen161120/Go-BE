package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "NDD")
	fmt.Println("----> My handler")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"users":   []string{"str1", "str2"},
	})
}
