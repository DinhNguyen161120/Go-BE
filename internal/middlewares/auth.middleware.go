package middlewares

import (
	"fmt"

	"example.com/m/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Println(token, " -- token")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
