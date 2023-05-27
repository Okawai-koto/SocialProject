package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Use for the get all users from mongo db database.
func UpdateUserMongo(userEmail string, body bson.M) bool {

	filter := bson.D{{"email", userEmail}}
	collection := getCollectionUsers()

	// MongoDB güncelleme ifadesini oluşturun
	update := bson.M{
		"$set": body,
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		println(err)
	}
	fmt.Println(result.ModifiedCount)
	if result.ModifiedCount == 0 {
		return false
	}
	return true
}
