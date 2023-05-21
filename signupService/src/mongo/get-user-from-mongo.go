package mongo

import (
	"context"
	"signupservice/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Use for the get all users from mongo db database.
func GetUserFromMongo(userEmail string) models.User {
	filter := bson.D{{"email", userEmail}}
	client := getMongoDBGetInstance()
	var result models.User
	err := client.Database("deneme").Collection("Users").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		println("çekme başarısız")
	}

	return result
}
