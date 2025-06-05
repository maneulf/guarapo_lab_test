package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maneulf/guarapo_lab_test/internal/handlers"
)

var tokenSplit = 7

type AuthMiddleware struct {
	*handlers.Base
}

func New(base *handlers.Base) *AuthMiddleware {
	return &AuthMiddleware{
		Base: base,
	}

}

func (m *AuthMiddleware) Auth(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if len(token) <= 7 {
		log.Println("Invalid token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized",
		})
		c.Abort()
		return
	}

	token = token[tokenSplit:]
	_, ok := m.Usernames[token]

	if !ok {
		log.Println("Invalid token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized",
		})
		c.Abort()
		return
	}
}
