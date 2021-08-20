package model

import (
	"encoding/json"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Topic struct {
	// Basic information
	TopicID   int64  `json:"topicId" bson:"topicId"`
	Name      string `json:"name", bson:"name"`
	SubjectID int64  `json:"subjectId" bson:"subjectId"`
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
