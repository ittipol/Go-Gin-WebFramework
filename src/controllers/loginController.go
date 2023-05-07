package controllers

import (
	"net/http"
	"web-api/models/request"
	"web-api/services/login"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
type loginController struct {
	loginService login.LoginService
}

func NewLoginController(loginService login.LoginService) LoginController {
	return &loginController{loginService}
}

func (obj *loginController) Login(c *gin.Context) {

	var body request.LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	res, err := obj.loginService.Login(body.Email, body.Password)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	// Set Response Header
	// c.Header("key", "value")

	c.SetCookie("refresh_token", res.RefreshToken, 1800, "/", "localhost", false, true)
	c.JSON(http.StatusOK, res)

	return
}

func (obj *loginController) Register(c *gin.Context) {

	var body request.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := obj.loginService.Register(body.Email, body.Password, body.Name)

	if err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

	return
}
