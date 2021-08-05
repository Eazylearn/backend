package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateTest(test model.Test) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}
	_, err := model.TestDB.Collection.InsertOne(context.TODO(), test)
	if err != nil {
		return err
	}
	return nil
}

func GetTestByID(id int64) (model.Test, error) {
	var test model.Test
	result, qErr := model.TestDB.Collection.Find(context.TODO(), bson.M{"testID": id})
	if qErr != nil {
		log.Println("test_repo.go/GetTestByID: Error finding", qErr.Error())
		return test, qErr
	}
	err := result.All(context.TODO(), &test)
	if err != nil {
		log.Println("test_repo.go/GetTestByID: Error encoding", err.Error())
		return test, err
	}
	return test, nil
}
