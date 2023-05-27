package mongo

import (
	"context"
	"signupservice/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Use for the get all users from mongo db database.
func GetUserFromMongo(userEmail string) models.User {
	filter := bson.D{{"email", userEmail}}
	collection := getCollectionUsers()
	var result models.User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		println("çekme başarısız")
	}

	return result
}
