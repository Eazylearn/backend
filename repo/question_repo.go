package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateQuestion(question model.Question) error {
	// return &Question{questionId: 1, name: "New topic", content: "abc", answer: "abc", testId: 1}
	_, err := model.QuestionDB.Collection.InsertOne(context.TODO(), question)
	if err != nil {
		return err
	}
	return nil
}
func GetAllQuestion() ([]model.Question, error) {
	list := make([]model.Question, 0)
	result, err := model.QuestionDB.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("question_repo/GetAllQuestion: ", err.Error())
		return list, err
	}
	//fmt.Println(result)
	result.All(context.TODO(), &list)
	//fmt.Println(list)
	return list, nil
}
func GetAllQuestionByTopicId(topicId string) ([]model.Question, error) {
	list := make([]model.Question, 0)
	result, err := model.QuestionDB.Collection.Find(context.TODO(), bson.M{"TopicId": topicId})

	if err != nil {
		log.Println("question_repo/GetAllQuestionByTopicId: ", err.Error())
		return list, err
	}
	result.All(context.TODO(), &list)
	return list, nil
}
func GetQuestioByIndex(index string) ([]model.Question, error) {
	list := make([]model.Question, 0)
	result := model.QuestionDB.Collection.FindOne(context.TODO(), bson.M{"Index": index})
	var question model.Question
	err := result.Decode(&question)
	if err != nil {
		log.Println("question_repo/GetQuestioByIndex: ", err.Error())
		return list, err
	}
	list = append(list, question)
	return list, nil
}
