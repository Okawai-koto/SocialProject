package routes

import (
	"net/http"
	"signupservice/src/mongo"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetUsers(c *gin.Context) {

	users := mongo.GetUsersFromMongo()
	c.IndentedJSON(http.StatusOK, users)
}
