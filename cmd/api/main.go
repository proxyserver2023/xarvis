package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Trainer struct {
	name string
	age  int
	city string
}

func main() {
	client, err := connectToMongoDB()
	checkErr(err)

	err = checkTheConnection(client)
	checkErr(err)

	collection := handleForCollection(client)

	dummyDataToAddInTheDatabase(collection)

	//updateADocument()
	//findSingleDocument()
	//deleteAllDocumentsInTheCollection()
	//closeTheConnection()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERRORR")
		log.Fatal(err)
	}
}

func connectToMongoDB() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)

	client, err := mongo.Connect(
		ctx,
		"mongodb://localhost:27017",
	)
	return client, err
}

func checkTheConnection(client *mongo.Client) error {
	ctx, _ := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	return client.Ping(ctx, nil)
}

func handleForCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("test").Collection("trainers")
}

func dummyDataToAddInTheDatabase(collection *mongo.Collection) {
	// Some dummy data to add to the Database
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	insertResult, err := collection.InsertOne(
		context.TODO(),
		ash,
	)
	checkErr(err)
	fmt.Println("Inserted a single Document: ", insertResult.InsertedID)

	// Insert multiple documents
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}
