package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Instance struct {
	CollectionName string
	DBName         string

	DB         *mongo.Database
	Collection *mongo.Collection
}