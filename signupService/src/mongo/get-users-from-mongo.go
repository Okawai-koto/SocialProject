package mongo

import (
	"context"
	"log"
	"signupservice/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Use for the get all users from mongo db database.
func GetUsersFromMongo() []models.User {
	var users = []models.User{}
	collection := getCollectionUsers()
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Belge dizisini doldurun
	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}
