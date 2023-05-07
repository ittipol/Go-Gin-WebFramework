package controllers

import (
	"net/http"
	"web-api/services/user"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Me(c *gin.Context)
}

type userController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) UserController {
	return &userController{userService}
}

func (obj *userController) Me(c *gin.Context) {

	id := c.MustGet("userId").(float64)

	user, err := obj.userService.Me(int(id))

	if err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
