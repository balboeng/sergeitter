package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = dbConnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://sbedoya:breve1213.@cluster0.rqdgp.mongodb.net/sergeitter?retryWrites=true&w=majority")

/*ConectarBD is a function that allow and cehck db connection */
func dbConnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	statusDb := client.Ping(context.TODO(), nil)
	if statusDb != nil {
		log.Fatal(statusDb.Error())
		return client
	}
	log.Println("<<--- SUCCESS DB CONNECTION ----->>")
	return client
}

/*CheckConnection checks ping connection to db*/
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
