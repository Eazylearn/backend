package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
)

func CreateQuestion(question model.Question) error {
	// return &Question{questionId: 1, name: "New topic", content: "abc", answer: "abc", testId: 1}
	_, err := model.QuestionDB.Collection.InsertOne(context.TODO(), question)
	if err != nil {
		return err
	}
	return nil
}