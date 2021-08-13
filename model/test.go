package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	// Basic information of test
	TestID int64  `json:"testId,omitempty" bson:"testId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Nnam,omitempty"`

	// Relative information of test
	TotalQuestion int32 `json:"totalQuestion,omitempty"`

	// Foreign keys
	//TopicID int64 `json:"topicId,omitempty"` // Reference to topic.go
	Questions []int32 `json:"questions,omitempty"`
}

var TestDB = &db.Instance{
	CollectionName: "test",
}

func InitTestDB(db *mongo.Database) {
	TestDB.ApplyDatabase(db)
}
