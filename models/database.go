package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/beegons/utils"
)

var mClient *mongo.Client
var mDatabase *mongo.Database
var db string

func ConnectToDB() {
	connect()
	ping()
}

func connect() {
	uri := utils.GetDBURL()

	var err error
	ctx, timeout := context.WithTimeout(context.Background(), 3*time.Second)
	defer timeout()
	mClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Connecting to MongoDB...")
}

func ping() {
	err := mClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection OK")
}

func getDatabase(dbName string) *mongo.Database {
	return mClient.Database(dbName)
}

func GetAppDatabase() *mongo.Database {
	return getDatabase(utils.GetAppDBName())
}

func GetCygnusDatabase() *mongo.Database {
	return getDatabase(utils.GetCygnusDBName())
}
