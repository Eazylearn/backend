package repo

import (
	"context"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSubject(subject model.Subject) error {
	// return &Topic{subjectId: 1, Name: "New subject"}
	_, err := model.SubjectDB.Collection.InsertOne(context.TODO(), subject)
	if err != nil {
		log.Println("subject_repo/CreateSubject: ", err.Error())
		return err
	}
	return nil
}

func GetSubjectByID(id int64) (model.Subject, error) {
	var subject model.Subject
	err := model.SubjectDB.Collection.FindOne(context.TODO(), bson.M{"subjectId": id}).Decode(&subject)
	if err != nil {
		log.Println("subject_repo/GetSubjectByID: ", err.Error())
		return subject, err
	}
	return subject, nil
}

func GetAllSubject() ([]model.Subject, error) {
	list := make([]model.Subject, 0)
	result, err := model.SubjectDB.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("subject_repo/GetAllSubject: ", err.Error())
		return list, err
	}
	result.All(context.TODO(), &list)
	return list, nil
}
