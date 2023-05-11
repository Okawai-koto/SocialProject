package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(databaseName string, collectionName string) (*mongo.Collection, error) {
	// MongoDB bağlantı seçenekleri
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://mongo-user:mongo-password@socialproject.uhpspae.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// MongoDB istemcisini oluştur
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// MongoDB istemcisini kullanarak koleksiyon nesnesi oluştur

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection := client.Database(databaseName).Collection(collectionName)
	return collection, nil

	// // Use the SetServerAPIOptions() method to set the Stable API version to 1
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI("mongodb+srv://mongo-user:mongo-password@socialproject.uhpspae.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// // Create a new client and connect to the server
	// client, err := mongo.Connect(context.TODO(), opts)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// // Send a ping to confirm a successful connection
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func AddUser(user, *mongo.collection) {
	// Kullanıcıyı "deneme" koleksiyonuna ekleyelim
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Eklenen kullanıcının ID'si: %v\n", res.InsertedID)
}
