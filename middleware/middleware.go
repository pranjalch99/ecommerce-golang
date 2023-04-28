package middleware

import (
	"fmt"
	"net/http"

	token "github.com/pranjalch99/ecommerce-golang/tokens"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	fmt.Println("Pranjal Authentication function is about to be executed")
	return func(c *gin.Context) {
		fmt.Println("Pranjal the client token is getting fetched")
		ClientToken := c.Request.Header.Get("token")
		fmt.Println("Pranjal the client token is ", ClientToken)
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
