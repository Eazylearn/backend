package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostTest struct {
	// Basic information of test
	TestID int64  `json:"testId,omitempty" bson:"testId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Nnam,omitempty"`

	// Relative information of test
	TotalQuestion int32  `json:"totalQuestion,omitempty"`
	Subject       string `json:"Subject,omitempty"  bson:"Subject,omitempty"`
	Type          string `json:"Type,omitempty"  bson:"Type,omitempty"` //combination v 1 subject
	// Foreign keys
	//TopicID int64 `json:"topicId,omitempty"` // Reference to topic.go
	Questions []int32 `json:"questions,omitempty"`
}
type Test struct {
	// Basic information of test
	TestID int64  `json:"testId,omitempty" bson:"testId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Nnam,omitempty"`

	// Relative information of test
	TotalQuestion int32  `json:"totalQuestion,omitempty"  bson:"TotalQuestion,omitempty"`
	Subject       string `json:"Subject,omitempty"  bson:"Subject,omitempty"`
	Type          string `json:"Type,omitempty"  bson:"Type,omitempty"` //combination v 1 subject
	// Foreign keys
	//TopicID int64 `json:"topicId,omitempty"` // Reference to topic.go
	Questions []Question `json:"questions,omitempty"  bson:"Questions,omitempty"`
}

var TestDB = &db.Instance{
	CollectionName: "test",
}

func InitTestDB(db *mongo.Database) {
	TestDB.ApplyDatabase(db)
}
