package middleware

import (
	"net/http"

	token "github.com/erdincmutlu/ecommerce/tokens"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no authorization header provided"})
			c.Abort()
			return
		}

		claims, errMsg := token.ValidateToken(ClientToken)
		if errMsg != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
