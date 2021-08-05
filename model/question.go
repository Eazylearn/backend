package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Question struct {
	// Basic information
	QuestionID int64 `json:"QuestionID,omitempty" bson:"QuestionID,omitempty"`

	// Relative information
	Content string `json:"Content,omitempty" bson:"Content,omitempty"`
	Answer  string `json:"Answer,omitempty" bson:"Answer,omitempty"`
	Reason  string `json:"Reason,omitempty" bson:"Reason,omitempty"`
	Type    string `json:"Type,omitempty" bson:"Type,omitempty" `
	Topic   string `json:"Topic,omitempty" bson:"Topic,omitempty" `
	Index   string `json:"Index,omitempty" bson:"Index,omitempty"`
	Subject string `json:"Subject,omitempty" bson:"Subject,omitempty"`
	Choices string `json:"Choices,omitempty" bson:"Choices,omitempty"`

	// Foreign keys
	TestID      string `json:"TestID,omitempty" bson:"TestID,omitempty" `
	Requirement string `json:"Requirement,omitempty" bson:"Requirement,omitempty"`
}

var QuestionDB = &db.Instance{
	CollectionName: "question",
}

func InitQuestionDB(db *mongo.Database) {
	QuestionDB.ApplyDatabase(db)
}
