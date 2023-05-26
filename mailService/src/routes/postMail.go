package routes

import (
	"mailservice/src/aws"

	"mailservice/src/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func PostMail(c *gin.Context) {

	var NewUser models.User
	if err := c.BindJSON(&NewUser); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, NewUser)

	code := verifycode()

	aws.SendEmail(NewUser.Email, code)

}
