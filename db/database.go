package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBContext context.Context

func getDBContext() context.Context {
	if DBContext == nil {
		DBContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
	}
	return DBContext
}

// Create only 1 table connect with database
func (i *Instance) ApplyDatabase(db *mongo.Database) *Instance {
	i.DB = db
	i.Collection = db.Collection(i.CollectionName)
	i.DBName = db.Name()
	return i
}

// Create only 1 database
func CreateUniversalDB(uri string, db string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Connect(getDBContext())
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(db)
}
