package db

import "go.mongodb.org/mongo-driver/mongo"

type ColsType struct {
	Users *mongo.Collection
}
