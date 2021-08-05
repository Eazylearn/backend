package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
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
	result, qErr := model.UserDB.Collection.Find(context.TODO(), bson.M{})
	if qErr != nil {
		log.Println("test_repo.go/GetTestByID: Error finding", qErr.Error())
		return test, qErr
	}
	fmt.Println(result == nil)
	err := result.All(context.TODO(), &test)
	if err != nil {
		log.Println("test_repo.go/GetTestByID: Error encoding", err.Error())
		return test, err
	}
	return test, nil
}
