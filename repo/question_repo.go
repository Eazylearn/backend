package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateQuestion(question []model.Question) error {
	// return &Question{questionId: 1, name: "New topic", content: "abc", answer: "abc", testId: 1}
	for i := 0; i < len(question); i++ {
		_, err := model.QuestionDB.Collection.InsertOne(context.TODO(), question)
		if err != nil {
			return err
		}
	}
	return nil
}
func GetAllQuestionByQuery(query *model.GetQuestionRequest) ([]model.Question, error) {
	filter := bson.M{}
	if query.TopicId != "" {
		filter["TopicId"] = query.TopicId
	}
	if query.Index != 0 {
		filter["Index"] = query.Index
	}
	list := make([]model.Question, 0)
	result, err := model.QuestionDB.Collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("question_repo/GetAllQuestion: ", err.Error())
		return list, err
	}
	//fmt.Println(result)
	result.All(context.TODO(), &list)
	//fmt.Println(list)
	return list, nil

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
func GetQuestioByIndex(index string) (model.Question, error) {
	var question model.Question
	result, qErr := model.QuestionDB.Collection.Find(context.TODO(), bson.M{"Index": index})

	if qErr != nil {
		log.Println("question_repo/GetQuestioByIndex: ", qErr.Error())
		return question, qErr
	}
	err := result.All(context.TODO(), &question)
	if err != nil {
		log.Println("question_repo.go/GetQuestioByIndex: Error encoding", err.Error())
		return question, err
	}

	return question, nil
}
