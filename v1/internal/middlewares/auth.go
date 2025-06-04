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

	//bodyBytes, err := io.ReadAll(c.Request.Body)

	//if err != nil {
	//log.Printf("Could not read json body, Err: %s", err)
	//}

	//c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	//var tasksRequest req.Task

	//err = json.Unmarshal(bodyBytes, &tasksRequest)

	//if err != nil {
	//log.Printf("Error could not unmarshal json body, Err: %s", err)
	//}

	//if err != nil {
	//log.Printf("Could not read json body, Err: %s", err)
	//c.Abort()
	//}

	token := c.GetHeader("Authorization")[tokenSplit:]

	_, ok := m.Usernames[token]

	if !ok {
		log.Println("Invalid token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized",
		})
		c.Abort()
		return
	}

	/*
		if strings.Compare(username, tasksRequest.Owner) != 0 {
			log.Println("Token does not match owner")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
			})
			c.Abort()
			return
		}
	*/
}
