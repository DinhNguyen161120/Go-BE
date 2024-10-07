package controller

import (
	"net/http"

	"example.com/m/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> servcie -> repo -> models -> dbs
func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUserService(),
		"users":   []string{"str1", "str2"},
		"userId":  "2asgqyueq6eq8q8qweyu",
	})
}
