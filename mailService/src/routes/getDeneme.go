package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetDeneme(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "it works (update mail)"})
}
