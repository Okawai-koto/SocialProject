package main

import (
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

var exampleuser = user{Email: "2", Password: "pasuwado"}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")

}
