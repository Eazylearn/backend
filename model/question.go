package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Question struct {
	// Basic information
	QuestionID int64  `json:"questionId"`
	Name       string `json:"name,omitempty"`

	// Relative information
	Content     string    `json:"content,omitempty" bson:"content,omitempty"`
	Requirement string    `json:"requirement" bson:"requirement"`
	Choices     [4]string `json:"choices,omitempty" bson:"choices,omitempty"`
	Answer      string    `json:"answer"`
	Reason      string    `json:"reason,omitempty" bson:"reason,omitempty"`

	// Foreign keys
	TestID  string `json:"testId,omitempty" bson:"testId,omitempty"`
	TopicID string `json:"topicId,omitempty" bson:"topicId,omitempty"`
}

var QuestionDB = &db.Instance{
	CollectionName: "question",
}

func InitQuestionDB(db *mongo.Database) {
	QuestionDB.ApplyDatabase(db)
}
