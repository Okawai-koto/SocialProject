package main

import (
	"signupservice/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/users/:email", routes.GetUserByEmail)
	router.GET("/users", routes.GetUsers)
	router.POST("/users", routes.PostUsers)
	router.DELETE("/users/:email", routes.DeleteUserByEmail)

	router.Run("localhost:8080")

}
