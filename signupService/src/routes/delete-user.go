package routes

import (
	"net/http"
	"signupservice/src/mongo"

	"github.com/gin-gonic/gin"
)

func DeleteUserByEmail(c *gin.Context) {
	userEmail := c.Param("email")

	result := mongo.DeleteUserFromMongo(userEmail)
	if result == false {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "it didn't work (i am talking about delete func ofc)"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "it works (i am talking about delete func ofc)"})
}
