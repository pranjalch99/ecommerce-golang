package middleware

import (
	"net/http"

	token "github.com/pranjalch99/ecommerce-golang/tokens"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "No authorisation header provided"})
			c.Abort()
			return
		}

		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
