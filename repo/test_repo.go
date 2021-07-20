package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
)

func CreateTest(test model.Test) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}
	_, err := model.TestDB.Collection.InsertOne(context.TODO(), test)
	if err != nil {
		return err
	}
	return nil
}