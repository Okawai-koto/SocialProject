package mongo

import (
	"context"
	"signupservice/models"
)

// Use for the add user to mongo db database.
func AddUserToMongo(user models.User) {
	client := GetMongoDBGetInstance()
	client.Database("deneme").Collection("Users").InsertOne(context.Background(), user)
}
