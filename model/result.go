package model

import (
	"time"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Result struct {
	// Basic information
	TimeStart time.Time `json:"timeStart,omitempty" bson:"timeStart,omitempty"`
	TimeEnd   time.Time `json:"timeEnd,omitempty" bson:"timeEnd,omitempty"`

	// Relative information
	Answer       [100]string `json:"answer,omitempty" bson:"answer,omitempty"`
	TotalCorrect int32       `json:"totalCorrect,omitempty" bson:"totalCorrect,omitempty"`
	//TotalTime    int64       `json:"totalTime" json:"totalTime"`

	// Foreign keys
	UserID int64 `json:"userId,omitempty" bson:"userId,omitempty"`
	TestID int64 `json:"testId,omitempty" bson:"testId,omitempty"`
}

var ResultDB = &db.Instance{
	CollectionName: "result",
}

func InitResultDB(db *mongo.Database) {
	ResultDB.ApplyDatabase(db)
}
