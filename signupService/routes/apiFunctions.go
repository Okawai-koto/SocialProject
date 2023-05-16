package routes

import (
	"net/http"
	"signupservice/models"
	"signupservice/mongo"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetUsers(c *gin.Context) {

	users := mongo.GetUsersFromMongo()
	c.IndentedJSON(http.StatusOK, users)
}

// postAlbums adds an album from JSON received in the request body.
func PostUsers(c *gin.Context) {
	var NewUser models.User

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&NewUser); err != nil {
		return
	}

	// Add the new album to the slice.
	mongo.AddUserToMongo(NewUser)

	c.IndentedJSON(http.StatusCreated, NewUser)
}

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
