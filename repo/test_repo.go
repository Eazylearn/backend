package repo

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertToCollection(test model.Test, collection string) error {
	if collection == "English" {
		_, err := model.Test_ENGDB.Collection.InsertOne(context.TODO(), test)
		return err
		//return err
	}
	if collection == "Math" {
		_, err := model.Test_MTHDB.Collection.InsertOne(context.TODO(), test)
		return err
		//return err
	}
	if collection == "Physic" {
		_, err := model.Test_PHYDB.Collection.InsertOne(context.TODO(), test)
		return err
		//return err
	}
	fmt.Printf("test_repo.go/CreateTest: Error cannot find subject or collection name")
	return nil
	//_, err := model.Test_ENGDB.Collection.InsertOne(context.TODO(), test)
}
func FindInCollection(filter bson.M, collection string) (*mongo.Cursor, error) {
	if collection == "English" {
		return model.Test_ENGDB.Collection.Find(context.TODO(), filter)

		//return err
	}
	if collection == "Math" {
		return model.Test_MTHDB.Collection.Find(context.TODO(), filter)

		//return err
	}
	fmt.Printf("test_repo.go/FindInCollection: Error cannot find subject or collection name")
	return nil, nil
	//_, err := model.Test_ENGDB.Collection.InsertOne(context.TODO(), test)
}
func CreateTestByBE(test model.PostTest) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}
	var questionArray []int64 = test.Questions[0:]
	list := make([]model.Question, 0)
	//questions:= []model.Question{}
	for i := 0; i < len(questionArray); i++ {
		var a int64 = questionArray[i]
		questions, qErr := GetQuestioByIndex(a)
		if qErr != nil {

			log.Println("test_repo.go/CreateTest: Error finding QuestionID   "+strconv.FormatInt(a, 10), qErr.Error())
		}
		list = append(list, questions)
	}

	body := model.Test{
		TestID:        test.TestID,
		Name:          test.Name,
		TotalQuestion: test.TotalQuestion,
		Subject:       test.Subject,
		Questions:     list,
		TopicID:       test.TopicID,
		Level:         test.Level,
		Type:          test.Type,
	}
	err := InsertToCollection(body, body.Subject) //model.TestDB.Collection.InsertOne(context.TODO(), body)

	if err != nil {
		log.Println("test_repo.go/CreateTest: Error Inserting", err.Error())
		return err
	}
	return nil
}
func CreateTest(test model.Test) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}

	err := InsertToCollection(test, test.Subject) //model.TestDB.Collection.InsertOne(context.TODO(), body)

	if err != nil {
		log.Println("test_repo.go/CreateTest: Error Inserting", err.Error())
		return err
	}
	return nil
}
func GetAllTestByQuery(query *model.Test) ([]model.Test, error) {
	filter := bson.M{}
	if query.Subject != "" {
		filter["Subject"] = query.Subject

	}
	if query.Name != "" {
		filter["Name"] = query.Name
	}
	if query.TestID != 0 {
		filter["TestID"] = query.TestID
	}
	if query.Level != 0 {
		filter["Level"] = query.Level
	}
	if len(query.TopicID)-1 != 0 {
		filter["TopicID"] = query.TopicID

	}

	list := make([]model.Test, 0)
	result, err := FindInCollection(filter, query.Subject)
	if err != nil {
		log.Println("test_repo/GetAllTestByQuery: error FindInCollection", err.Error())
		return list, err
	}
	result.All(context.TODO(), &list)
	return list, nil
}
