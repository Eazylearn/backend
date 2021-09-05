package model

import (
	"encoding/json"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Topic struct {
	// Basic information
	TopicID   string `json:"topicId" bson:"topicId"`
	Name      string `json:"name" bson:"name"`
	SubjectID string `json:"subjectId" bson:"subjectId"`
}

func (t Topic) String() string {
	tjson, _ := json.Marshal(t)
	return string(tjson)
}

var TopicDB = &db.Instance{
	CollectionName: "topic",
}

func InitTopicDB(db *mongo.Database) {
	TopicDB.ApplyDatabase(db)
}
