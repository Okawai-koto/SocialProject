package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// albums slice to seed record album data.
var users = []user{
	{Email: "1", Password: "Blue Train"},
	{Email: "2", Password: "Jeru"},
}

var exampleuser = user{Email: "1", Password: "Blue Train"}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
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
	users = append(users, newUser)
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

func main() {

	collection, err := GetCollection("deneme", "Users")
	if err != nil {
		// Hata oluştu
		fmt.Println(err)
		return
	}

	AddUser(exampleuser, collection)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")

}
