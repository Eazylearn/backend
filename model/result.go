package model

import (
	"time"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Result struct {
	// Basic information
	TimeStart time.Time `json:"timeStart" json:"timeStart"`
	TimeEnd   time.Time `json:"timeEnd" json:"timeEnd"`

	// Relative information
	Answer       [100]string `json:"answer" bson:"answer"`
	TotalCorrect int32       `json:"totalCorrect" bson:"totalCorrect"`
	TotalTime    int64       `json:"totalTime" json:"totalTime"`

	// Foreign keys
	UserID int64 `json:"userId"`
	TestID int64 `json:"testId"`
}

var ResultDB = &db.Instance{
	CollectionName: "result",
}

func InitResultDB(db *mongo.Database) {
	ResultDB.ApplyDatabase(db)
}
