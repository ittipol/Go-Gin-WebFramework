package controllers

import (
	"net/http"
	"web-api/models/request"
	"web-api/services/login"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginController struct {
	db *gorm.DB
}

type ILoginController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func NewLoginController(db *gorm.DB) ILoginController {
	return &loginController{
		db,
	}
}

func (h *loginController) Login(c *gin.Context) {

	var body request.LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	s := login.NewLoginService(h.db)
	res, err := s.Login(&body)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	// c.Header("Set-Cookie", "refresh_token="+res.RefreshToken)

	c.SetCookie("refresh_token", res.RefreshToken, 1800, "/", "localhost", false, true)
	c.JSON(http.StatusOK, res)

}

func (h *loginController) Register(c *gin.Context) {

	s := login.NewLoginService(h.db)
	err := s.Register(c)

	if err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
