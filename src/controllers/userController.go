package controllers

import (
	"net/http"
	"web-api/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

type IUserController interface {
	Me(c *gin.Context)
}

func NewUserController(db *gorm.DB) IUserController {
	return &userController{db}
}

func (h *userController) Me(c *gin.Context) {

	id := c.MustGet("userId").(float64)

	s := user.NewUserService(h.db)
	res, err := s.Me(int(id))

	if err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, res)
}
