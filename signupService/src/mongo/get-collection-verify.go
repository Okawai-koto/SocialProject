package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Use for the get mongo db client instance
func getCollectionVerify() *mongo.Collection {

	client := getMongoDBGetInstance()

	instanceVerifyCollection := client.Database("deneme").Collection("verify")

	return instanceVerifyCollection
}
