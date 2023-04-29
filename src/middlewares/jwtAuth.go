package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(s, "Bearer ")

	fmt.Println(tokenString)

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid Token",
		})

		return
	}

	hmacSampleSecret := []byte(os.Getenv("JWT_ACCESS_TOKEN"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("userId", claims["userId"])
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	}
}
