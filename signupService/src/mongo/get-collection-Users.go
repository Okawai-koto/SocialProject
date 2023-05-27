package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Use for the get mongo db client instance
func getCollectionUsers() *mongo.Collection {

	client := getMongoDBGetInstance()

	instanceUsersCollection := client.Database("deneme").Collection("Users")

	return instanceUsersCollection
}
