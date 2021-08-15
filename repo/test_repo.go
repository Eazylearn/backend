package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateTest(test model.PostTest) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}
	var questionArray []int32 = test.Questions[0:]
	list := make([]model.Question, 0)
	//questions:= []model.Question{}
	for i := 0; i < len(questionArray); i++ {
		questions, qErr := GetQuestioByIndex(string(questionArray[i]))
		if qErr != nil {
			log.Println("test_repo.go/CreateTest: Error finding QuestionID"+string(questionArray[i]), qErr.Error())
		}
		list = append(list, questions)
	}

	body := model.Test{
		TestID:        test.TestID,
		Name:          test.Name,
		TotalQuestion: test.TotalQuestion,
		Questions:     list,
	}
	_, err := model.TestDB.Collection.InsertOne(context.TODO(), body)

	if err != nil {
		log.Println("test_repo.go/CreateTest: Error Inserting", err.Error())
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
