package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	mser "./mog_server"
)



func main(){
	fmt.Println("start...")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("dbtestgo")

	//mser.AddUser(db)
	mser.FindUser(db)
	mser.FindManyUser(db)
	mser.FindOneUser(db)



	for{ }

	/*err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}*/
	fmt.Println("Connection to MongoDB closed.")
	fmt.Println("end.")
}