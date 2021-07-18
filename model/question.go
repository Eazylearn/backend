package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Question struct {
	// Basic information
	QuestionID int64  `json:"questionId,omitempty"`
	Name       string `json:"name,omitempty"`

	// Relative information
	Content string `json:"content,omitempty"`
	Answer  string `json:"answer,omitempty"`

	// Foreign keys
	TestID string `json:"testId,omitempty"`
}

var QuestionDB = &db.Instance{
	CollectionName: "question",
}

func InitQuestionDB(db *mongo.Database) {
	QuestionDB.ApplyDatabase(db)
}