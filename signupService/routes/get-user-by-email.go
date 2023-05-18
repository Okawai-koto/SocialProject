package routes

import (
	"fmt"
	"net/http"
	"signupservice/mongo"

	"github.com/gin-gonic/gin"
)

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetUserByEmail(c *gin.Context) {
	userEmail := c.Param("email")
	result := mongo.GetUserFromMongo(userEmail)
	fmt.Println(result)

	if result.Email == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "email not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}
