package model

import (
	"time"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Result struct {
	// Basic information
	ResultID  int64      `json:"resultId,omitempty"`
	TimeStart *time.Time `json:"timeStart,omitempty"`
	TimeEnd   *time.Time `json:"timeEnd,omitempty"`

	// Relative information
	TotalCorrect     int32      `json:"totalCorrect,omitempty"`
	TotalIncorrect   int32      `json:"totalIncorrect,omitempty"`
	AnsweredQuestion int32      `json:"answeredQuestion,omitempty"`
	TotalTime        *time.Time `json:"totalTime,omitempty"`

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