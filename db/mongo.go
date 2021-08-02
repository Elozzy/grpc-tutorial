package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB            *mongo.Collection
	DBCollections ColsType
	MongoClient   *mongo.Client
	err           error
)

//ConnectMongo ...
func ConnectMongo() {
	to := time.Second * 60
	clientOpts := options.ClientOptions{ConnectTimeout: &to}
	clientOpts.SetDirect(true)

	runningENV := os.Getenv("APP_ENV")
	switch runningENV {
	case "production":
		clientOpts.ApplyURI(os.Getenv("PRODUCTION_CONN"))
		log.Println("Running on production")
	case "sit":
		clientOpts.ApplyURI(os.Getenv("SIT_CONN"))
		log.Println("Running on development [SIT DB]..")
	case "dev":
		clientOpts.ApplyURI(os.Getenv("DEVELOPMENT_CONN"))
		log.Println("Running on development [LOCAL DB]..")

	}

	MongoClient, err = mongo.Connect(context.TODO(), &clientOpts)
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}

	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongoDB.")

	DBCollections.Users = MongoClient.Database("grpc").Collection("users")

}

//GetMongoDB ...
func GetMongoDB() *ColsType {
	return &DBCollections
}
