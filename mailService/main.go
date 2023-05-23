package main

import (
	"mailservice/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// router.GET("/users/:email", routes.GetUserByEmail)
	router.GET("/deneme", routes.GetDeneme)
	// router.POST("/users", routes.PostUsers)
	// router.PATCH("/users/:email", routes.UpdateUserEmail)
	// router.DELETE("/users/:email", routes.DeleteUserByEmail)

	router.Run(":8080")

}
