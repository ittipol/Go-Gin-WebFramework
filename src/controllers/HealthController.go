package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Health(c *gin.Context)
}

type healthController struct {
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (healthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
