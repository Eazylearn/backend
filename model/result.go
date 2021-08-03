package model

import (
	"time"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Result struct {
	// Basic information
	ResultID  int64      `json:"resultId" bson:"resultId"`
	TimeStart *time.Time `json:"timeStart,omitempty" json:"timeStart,omitempty"`
	TimeEnd   *time.Time `json:"timeEnd,omitempty" json:"timeEnd,omitempty"`

	// Relative information
	Answer           [1000]string `json:"answer,omitempty" bson:"answer,omitempty"`
	TotalCorrect     int32        `json:"totalCorrect,omitempty" bson:"totalCorrect,omitempty"`
	TotalIncorrect   int32        `json:"totalIncorrect,omitempty" bson:"totalIncorrect,omitempty"`
	AnsweredQuestion int32        `json:"answeredQuestion,omitempty" bson:"answeredQuestion,omitempty"`
	TotalTime        *time.Time   `json:"totalTime,omitempty" json:"totalTime,omitempty"`

	// Foreign keys
	UserID int64 `json:"userId,omitempty"`
	TestID int64 `json:"testId,omitempty"`
}

var ResultDB = &db.Instance{
	CollectionName: "result",
}

func InitResultDB(db *mongo.Database) {
	ResultDB.ApplyDatabase(db)
}
