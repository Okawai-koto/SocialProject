package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Use for the get all users from mongo db database.
func DeleteUserFromMongo(userEmail string) bool {
	filter := bson.D{{"email", userEmail}}
	collection := getCollectionUsers()

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		println(err)
	}
	if result.DeletedCount == 0 {
		return false
	}
	return true
}
