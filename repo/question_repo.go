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
	if query.Level != 0 {
		filter["Level"] = query.Level
	}
	if query.SubjectId != "" {
		filter["SubjectId"] = query.SubjectId
	}

	list := make([]model.Question, 0)
	result, err := model.QuestionDB.Collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("question_repo/GetAllQuestion: ", err.Error())
		return list, err
	}

	//fmt.Println(result)
	result.All(context.TODO(), &list)
	/*if len(list) == 0 {
		log.Println("question_repo/GetAllQuestion: wrong parameter", err.Error())
		return list, err
	}*/
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
func GetQuestioByIndex(index int64) (model.Question, error) {
	var question model.Question
	err := model.QuestionDB.Collection.FindOne(context.TODO(), bson.M{"Index": index}).Decode(&question)

	if err != nil {
		log.Println("question_repo/GetQuestioByIndex: ", err.Error())
		return question, err
	}

	return question, nil
}
