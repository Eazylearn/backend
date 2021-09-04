package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Question struct {
	// Basic information
	//QuestionId int64  `json:"1uestionId" bson:"QuestionId"`
	Name string `json:"name,omitempty" bson:"Name,omitempty"`

	Content     string `json:"content,omitempty" bson:"Content,omitempty"`
	Answer      string `json:"cnswer,omitempty" bson:"Answer,omitempty"`
	Reason      string `json:"reason,omitempty" bson:"Reason,omitempty"`
	Type        string `json:"type,omitempty" bson:"Type,omitempty" `
	Requirement string `json:"requirement,omitempty" bson:"Requirement,omitempty"`
	Index       int64  `json:"index,omitempty" bson:"Index,omitempty"`
	Subject     string `json:"subject,omitempty" bson:"Subject,omitempty"`
	Choices     string `json:"choices,omitempty" bson:"Choices,omitempty"`
	// Foreign keys
	TestId  string `json:"testId,omitempty" bson:"TestId,omitempty" `
	TopicId string `json:"topicId,omitempty" bson:"TopicId,omitempty" `
	Level   int64  `json:"level,omitempty" bson:"Level,omitempty" `

	//questionId int64  `json:"questionId bson:"questionId"`
	//	name       string `json:"name,omitempty bson:"questionId,omitempty"`

	// Relative information
	/*content     string    `json:"content,omitempty" bson:"content,omitempty"`
	requirement string    `json:"requirement,omitempty" bson:"requirement,omitempty"`
	choices     [4]string `json:"choices,omitempty" bson:"choices,omitempty"`
	answer      string    `json:"answer,omitempty  bson:"answer,omitempty"`
	reason      string    `json:"reason,omitempty" bson:"reason,omitempty"`
	qType       string    `json:"qType,omitempty" bson:"qType,omitempty" `
	index       string    `json:"index,omitempty" bson:"index,omitempty"`
	subject     string    `json:"subject,omitempty" bson:"subject,omitempty"`

	// Foreign keys
	testId  string `json:"testId,omitempty" bson:"testId,omitempty"`
	topicId string `json:"topicId,omitempty" bson:"topicId,omitempty"`*/
}

type GetQuestionRequest struct {
	Question
}

var QuestionDB = &db.Instance{
	CollectionName: "question",
}

func InitQuestionDB(db *mongo.Database) {
	QuestionDB.ApplyDatabase(db)
}
