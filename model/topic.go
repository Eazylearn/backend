package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Topic struct {
	// Basic information
	TopicID   int64  `json:"topicId" bson:"topicId"`
	Name      string `json:"name" bson:"name"`
	SubjectID int64  `json:"subjectId" bson:"subjectId"`
}

var TopicDB = &db.Instance{
	CollectionName: "topic",
}

func InitTopicDB(db *mongo.Database) {
	TopicDB.ApplyDatabase(db)
}
