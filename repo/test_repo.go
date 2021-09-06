package repo

import (
	"context"
	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"sort"
	"strconv"
	"strings"
	//	"go.mongodb.org/mongo-driver/mongo"
)

/*func InsertToCollection(test model.Test, collection string) error {
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
*/
func CreateTestByBE(test []model.PostTest) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}
	for i := 0; i < len(test); i++ {
		var questionArray []int64 = test[i].Questions[0:]
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
			TestId:        test[i].TestId,
			Name:          test[i].Name,
			TotalQuestion: test[i].TotalQuestion,
			Subject:       test[i].Subject,
			Questions:     list,
			TopicId:       test[i].TopicId,
			Level:         test[i].Level,
			Type:          test[i].Type,
			SubjectId:     test[i].SubjectId,
		}
		sort.Strings(body.TopicId)
		//err := InsertToCollection(body, body.Subject) //

		_, err := model.TestDB.Collection.InsertOne(context.TODO(), body)
		if err != nil {
			log.Println("test_repo.go/CreateTest: Error Inserting", err.Error())
			return err
		}

	}
	return nil
}
func CreateTest(test model.Test) error {
	// return &Test{testId: 1, Name: "New test", totalQuestion: 0, topicId: 1}

	//err := InsertToCollection(test, test.Subject) //
	sort.Strings(test.TopicId)
	_, err := model.TestDB.Collection.InsertOne(context.TODO(), test)

	if err != nil {
		log.Println("test_repo.go/CreateTest: Error Inserting", err.Error())
		return err
	}
	return nil
}
func RemoveIndex(s []model.Test, index int) []model.Test {
	return append(s[:index], s[index+1:]...)
}
func GetAllTestByQuery(query *model.Test) ([]model.Test, error) {
	filter := bson.M{}
	list := make([]model.Test, 0)
	if query.Subject == "" && query.SubjectId == "" && len(query.TopicId) != 0 {
		return list, nil
	}
	if query.Subject != "" {
		filter["Subject"] = query.Subject

	}
	if query.SubjectId != "" {
		filter["SubjectId"] = query.SubjectId

	}
	if query.Name != "" {
		filter["Name"] = query.Name
	}
	if query.TestId != 0 {
		filter["TestId"] = query.TestId
	}
	if query.Level != 0 {
		filter["Level"] = query.Level
	}
	if len(query.TopicId) == 1 {
		filter["TopicId"] = query.TopicId

	}

	//list := make([]model.Test, 0)
	//result, err := FindInCollection(filter, query.Subject)
	result, err := model.TestDB.Collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("test_repo/GetAllTestByQuery: error FindInCollection", err.Error())
		return list, err
	}

	result.All(context.TODO(), &list)
	if len(query.TopicId) > 1 {

		//sort.Strings(query.TopicId)
		for i := 0; i < len(query.TopicId); i++ {
			var a string = query.TopicId[i]
			filter[a] = a
		}
		for i := 0; i < len(list); i++ {
			var topics []string = list[i].TopicId
			for j := 0; j < len(topics); j++ {
				if filter[topics[j]] != topics[j] {
					list = RemoveIndex(list, i)
					continue
				}
			}
		}

	}
	return list, nil
}
func GetTestTotalQuestion(TestID int64) (int32, error) {
	var test model.Test
	err := model.TestDB.Collection.FindOne(context.TODO(), bson.M{"TestId": TestID}).Decode(&test)

	if err != nil {
		log.Println("test_repo/GetTestTotalQuestion: Cannot find test id", err.Error())
		return 0, err
	}

	return int32(test.TotalQuestion), nil
}
func GetTestTotalCorrect(TestID int64, Answers []string) (int32, error) {
	var test model.Test
	err := model.TestDB.Collection.FindOne(context.TODO(), bson.M{"TestId": TestID}).Decode(&test)

	if err != nil {
		log.Println("test_repo/GetTestTotalCorrect: Cannot find test id", err.Error())
		return 0, err
	}
	var questions []model.Question = test.Questions
	var score int32 = 0
	for i := 0; i < len(questions); i++ {
		var qAnswer string = strings.TrimSpace(questions[i].Answer)
		if qAnswer == strings.TrimSpace(Answers[i]) {
			score++
		}

	}
	return score, nil

}
