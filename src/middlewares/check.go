package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"Error": "",
	})

}
