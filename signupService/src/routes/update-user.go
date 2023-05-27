package routes

import (
	"net/http"
	"signupservice/src/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// getAlbums responds with the list of all albums as JSON.
func UpdateUserEmail(c *gin.Context) {

	// İstek gövdesini JSON olarak alın
	var requestBody bson.M
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz istek gövdesi"})
		return
	}
	userEmail := c.Param("email")
	result := mongo.UpdateUserMongo(userEmail, requestBody)
	if result == false {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "it didn't work (uptade mail)"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "it works (update mail)"})
}
