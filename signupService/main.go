package main

import (
	"signupservice/routes"

	"github.com/gin-gonic/gin"
)

// user represents data about a record album.
type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	router := gin.Default()
	router.GET("/users", routes.GetUsers)
	router.POST("/users", routes.PostUsers)

	router.GET("/users/:email", routes.GetUserByEmail)

	router.Run("localhost:8080")

}
