package model

import (

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Topic struct {
	// Basic information
	TopicID int64  `json:"topicId,omitempty"`
	Name    string `json:"name,omitempty"`
}

var TopicDB = &db.Instance{
	CollectionName: "topic",
}

func InitTopicDB(db *mongo.Database) {
	TopicDB.ApplyDatabase(db)
}