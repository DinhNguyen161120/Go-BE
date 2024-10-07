package controller

import (
	"example.com/m/internal/service"
	"example.com/m/pkg/response"
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
	if false {
		response.ErrorResponse(c, 20003, "No need")
	}
	response.SuccessResponse(c, 20001, []string{"data1", "data2", "data3"})
}
