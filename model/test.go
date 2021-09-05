package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostTest struct {
	// Basic information of test
	TestId int64  `json:"testId,omitempty" bson:"testId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Name,omitempty"`

	// Relative information of test
	TotalQuestion int64  `json:"totalQuestion,omitempty"`
	Subject       string `json:"subject,omitempty"  bson:"Subject,omitempty"`
	Type          string `json:"type,omitempty"  bson:"Type,omitempty"` //combination v 1 subject
	Level         int64  `json:"level,omitempty"  bson:"Level,omitempty"`
	// Foreign keys
	TopicId   []string `json:"topicId,omitempty"   bson:"TopicId"` // Reference to topic.go
	Questions []int64  `json:"questions,omitempty"`
}
type Test struct {
	// Basic information of test
	TestId int64  `json:"testId,omitempty" bson:"TestId,omitempty"`
	Name   string `json:"name,omitempty" bson:"Name,omitempty"`

	// Relative information of test
	TotalQuestion int64  `json:"totalQuestion,omitempty"  bson:"TotalQuestion,omitempty"`
	Subject       string `json:"subject,omitempty"  bson:"Subject,omitempty"`
	Type          string `json:"type,omitempty"  bson:"Type,omitempty"` //combination v 1 subject
	Level         int64  `json:"level,omitempty"  bson:"Level,omitempty"`
	// Foreign keys
	TopicId   []string   `json:"topicId,omitempty"   bson:"TopicId"` // Reference to topic.go
	Questions []Question `json:"questions,omitempty"  bson:"Questions,omitempty"`
}

/*var Test_ENGDB = &db.Instance{
	CollectionName: "test_ENG",
}
var Test_MTHDB = &db.Instance{
	CollectionName: "test_MTH",
}
var Test_PHYDB = &db.Instance{
	CollectionName: "test_PHYDB",
}*/
var TestDB = &db.Instance{
	CollectionName: "test",
}

func InitTestDB(db *mongo.Database) {
	//Test_ENGDB.ApplyDatabase(db)
	//Test_MTHDB.ApplyDatabase(db)
	//Test_PHYDB.ApplyDatabase(db)
	TestDB.ApplyDatabase(db)
}
