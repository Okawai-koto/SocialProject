package routes

import (
	"mailservice/src/aws"

	"mailservice/src/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func PostMail(c *gin.Context) {

	var UsersEmailData models.MailData
	if err := c.BindJSON(&UsersEmailData); err != nil {
		return
	}

	result := aws.SendEmail(UsersEmailData.Email, UsersEmailData.Verifycode)

	if result == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "E mail couldnt send"})
	}
	c.IndentedJSON(http.StatusOK, UsersEmailData)
}
