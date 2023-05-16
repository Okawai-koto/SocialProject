package mongo

import (
	"context"
	"fmt"
	"log"
	"signupservice/models"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// It is necessary for object that we will create. We want to create just 1 object. U can search singleton model :)
var (
	once     sync.Once
	instance *mongo.Client
)

// Use for the get mongo db client instance
func getMongoDBGetInstance() *mongo.Client {
	once.Do(func() {
		// Use the SetServerAPIOptions() method to set the Stable API version to 1
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI("mongodb+srv://mongo-user:mongo-password@socialproject.uhpspae.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}

		// Send a ping to confirm a successful connection
		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
			panic(err)
		}

		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

		instance = client
	})

	return instance
}

// Use for the add user to mongo db database.
func AddUserToMongo(user models.User) {
	client := getMongoDBGetInstance()
	client.Database("deneme").Collection("Users").InsertOne(context.Background(), user)
}

// Use for the get all users from mongo db database.
func GetUsersFromMongo() []models.User {
	var users = []models.User{}
	client := getMongoDBGetInstance()
	cur, err := client.Database("deneme").Collection("Users").Find(context.TODO(), bson.D{})
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

// func getUserFromMongoByMail(userEmail string) user {

// 	client := getMongoDBGetInstance()
// 	filter := bson.D{{"email", userEmail}}
// 	var result user
// 	err := client.Database("deneme").Collection("Users").FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// This error means your query did not match any documents.
// 			return result
// 		}
// 		panic(err)
// 	}

// 	return result
// }
