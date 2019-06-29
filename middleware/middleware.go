package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorization(c* gin.Context){
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized,gin.H{"message":http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}
	c.Next()
}

