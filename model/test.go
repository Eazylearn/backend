package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	// Basic information of test
	TestID int64  `json:"testId,omitempty"`
	Name   string `json:"name,omitempty"`

	// Relative information of test
	TotalQuestion int32 `json:"totalQuestion,omitempty"`

	// Foreign keys
	TopicID int64 `json:"topicId,omitempty"` // Reference to topic.go
}

var TestDB = &db.Instance{
	CollectionName: "ha",
}

func InitTestDB(db *mongo.Database) {
	TestDB.ApplyDatabase(db)
}
