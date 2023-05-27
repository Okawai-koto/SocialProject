package mongo

import (
	"context"
	"signupservice/src/models"
)

// Use for the add user to mongo db database.
func AddUserToMongo(user models.User) {
	collection := getCollectionUsers()
	collection.InsertOne(context.Background(), user)
}
