package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Question struct {
	// Basic information
	questionId int64  `json:"questionId bson:"questionId"`
	name       string `json:"name,omitempty bson:"questionId,omitempty"`

	// Relative information
	content     string    `json:"content,omitempty" bson:"content,omitempty"`
	requirement string    `json:"requirement" bson:"requirement"`
	choices     [4]string `json:"choices,omitempty" bson:"choices,omitempty"`
	answer      string    `json:"answer  bson:"answer"`
	reason      string    `json:"reason,omitempty" bson:"reason,omitempty"`
	qType       string    `json:"qType,omitempty" bson:"qType,omitempty" `
	index       string    `json:"index,omitempty" bson:"index,omitempty"`
	subject     string    `json:"subject,omitempty" bson:"subject,omitempty"`

	// Foreign keys
	testId  string `json:"testId,omitempty" bson:"testId,omitempty"`
	topicId string `json:"topicId,omitempty" bson:"topicId,omitempty"`
}

var QuestionDB = &db.Instance{
	CollectionName: "question",
}

func InitQuestionDB(db *mongo.Database) {
	QuestionDB.ApplyDatabase(db)
}
