package mongo

import (
	"context"
	"fmt"
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

		// Çevresel değişkenin adını belirleyin
		// credential := "MONGO_URI"

		// Çevresel değişkeni okuyun
		// mongoUri := os.Getenv(credential)

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
		fmt.Println("This is from signupservice ver:0.0.8.1")

		instance = client
	})

	return instance
}
