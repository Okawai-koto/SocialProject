package main

import (
	"signupservice/aws"
	"signupservice/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	aws.SendEmail()
	router := gin.Default()

	router.GET("/users/:email", routes.GetUserByEmail)
	router.GET("/users", routes.GetUsers)
	router.POST("/users", routes.PostUsers)
	router.PATCH("/users/:email", routes.UpdateUserEmail)
	router.DELETE("/users/:email", routes.DeleteUserByEmail)

	router.Run("localhost:8080")

}
