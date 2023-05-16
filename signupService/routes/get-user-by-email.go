package routes

import (
	"net/http"
	"signupservice/mongo"

	"github.com/gin-gonic/gin"
)

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	users := mongo.GetUsersFromMongo()
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.Email == email {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "email not found"})
}
