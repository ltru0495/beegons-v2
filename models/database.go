package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mClient *mongo.Client
var mDatabase *mongo.Database
var db string

func ConnectToDB(host string, port int, database string) {
	connect(host, port)
	ping()
	db = database
}

func connect(host string, port int) {
	uri := fmt.Sprintf("mongodb://%s:%d/", host, port)

	var err error
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()
	mClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Connected to MongoDB...")
}

func ping() {
	err := mClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection OK")
}

func getDatabase() *mongo.Database {
	if mDatabase == nil {
		mDatabase = mClient.Database(db)
	}
	return mDatabase
}
