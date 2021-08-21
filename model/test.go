package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostTest struct {
	// Basic information of test
	TestID int64  `json:"testId,omitempty" bson:"testId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Name,omitempty"`

	// Relative information of test
	TotalQuestion int32  `json:"totalQuestion,omitempty"`
	Subject       string `json:"subject,omitempty"  bson:"Subject,omitempty"`
	Type          string `json:"type,omitempty"  bson:"Type,omitempty"` //combination v 1 subject
	// Foreign keys
	//TopicID int64 `json:"topicId,omitempty"` // Reference to topic.go
	Questions []int32 `json:"questions,omitempty"`
}
type Test struct {
	// Basic information of test
	TestID int64  `json:"testId,omitempty" bson:"TestId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Name,omitempty"`

	// Relative information of test
	TotalQuestion int32  `json:"totalQuestion,omitempty"  bson:"TotalQuestion,omitempty"`
	Subject       string `json:"subject,omitempty"  bson:"Subject,omitempty"`
	Type          string `json:"type,omitempty"  bson:"Type,omitempty"` //combination v 1 subject
	// Foreign keys
	//TopicID int64 `json:"topicId,omitempty"` // Reference to topic.go
	Questions []Question `json:"questions,omitempty"  bson:"Questions,omitempty"`
}

var Test_ENGDB = &db.Instance{
	CollectionName: "test_ENG",
}
var Test_MTHDB = &db.Instance{
	CollectionName: "test_MTH",
}

func InitTestDB(db *mongo.Database) {
	Test_ENGDB.ApplyDatabase(db)
	Test_MTHDB.ApplyDatabase(db)
}
