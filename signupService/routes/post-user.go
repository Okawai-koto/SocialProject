package routes

import (
	"net/http"
	"signupservice/models"
	"signupservice/mongo"

	"github.com/gin-gonic/gin"
)

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
