package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func AddUser(user user) {
	client := getMongoDBGetInstance()
	client.Database("deneme").Collection("Users").InsertOne(context.Background(), user)
}

func GetUsers() {
	client := getMongoDBGetInstance()
	cur, err := client.Database("deneme").Collection("Users").Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Belge dizisini doldurun
	for cur.Next(context.TODO()) {
		var user user
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	fmt.Print(users)
}