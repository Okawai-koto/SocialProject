package routes

import (
	"net/http"
	"signupservice/src/functions"
	"signupservice/src/models"
	"signupservice/src/mongo"

	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body.
func SignUpUser(c *gin.Context) {

	verifycode := functions.Verifycode()

	var NewUser models.User

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&NewUser); err != nil {
		return
	}

	verify := models.VerifyCodeTemplate{Email: NewUser.Email, VerifyCode: verifycode}

	mongo.AddUserToMongo(NewUser)

	mongo.AddVerifyToMongo(verify)

	functions.PostToSignUpService(NewUser.Email, verifycode)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "it works (signup worked)"})
}
