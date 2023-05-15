package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	GetUsers()
	c.IndentedJSON(http.StatusOK, users)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	AddUser(newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	email := c.Param("email")

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
