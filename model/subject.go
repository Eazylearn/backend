package model

import (
	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Subject struct {
	// Basic information
	SubjectID string `json:"subjectId" bson:"subjectId"`
	Name      string `json:"name" bson:"name"`
}

var SubjectDB = &db.Instance{
	CollectionName: "subject",
}

func InitSubjectDB(db *mongo.Database) {
	SubjectDB.ApplyDatabase(db)
}
