package repo

import (
	"context"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"math"
)

func CreateResult(result model.Result) error {
	//result.TotalCorrect = GetTestTotalCorrect(result.TestID, result.Answer)

	//temp.TotalCorrect,_  =GetTestTotalCorrect(result.TestID,result.Answer[:])
	temp := model.Result{
		TimeStart: result.TimeStart,
		TimeEnd:   result.TimeEnd,
		Answer:    result.Answer,
		//TotalCorrect: 0,
		UserID: result.UserID,
		TestID: result.TestID,
	}
	temp.TotalCorrect, _ = GetTestTotalCorrect(result.TestID, result.Answer[:])
	_, err := model.ResultDB.Collection.InsertOne(context.TODO(), temp)
	if err != nil {
		log.Println("result_repo/CreateResult: ", err.Error())
		return err
	}

	return nil
}

func GetResultByUserID(userId int64) ([]model.Result, error) {
	list := make([]model.Result, 0)
	result, err := model.ResultDB.Collection.Find(context.TODO(), bson.M{"userId": userId})
	//println(list[0].TestID)
	if err != nil {
		log.Println("result_repo/GetResultByUserID: error encoding ", err.Error())
		return list, err
	}

	result.All(context.TODO(), &list)

	return list, nil
}

func GetResultScore(result model.Result) float64 {

	totalQuestion, _ := GetTestTotalQuestion(result.TestID)
	if totalQuestion != 0 {
		return math.Round(float64(result.TotalCorrect)*10 / float64(totalQuestion))
	}
	//score = float64(result.TotalCorrect * 10 / totalQuestion)

	return -1
}

/*func GetUserHistoryResult(result model.Result) ([]model.Result, error) {
	listResult := make([]model.Result, 0)
	if result.TimeEnd.Before(result.TimeStart) {
		//log.Printf("1")
		return listResult, nil

		//log.Printf("1")
	}

	if (result.TimeEnd == time.Time{}) {
		result.TimeEnd = time.Now()
		//log.Printf("2")
	}
	listResult, rErr := GetResultByUserID(result.UserID)
	if rErr != nil {
		log.Println("result_repo/GetUserHistoryResult: error encoding result ", rErr.Error())
		return listResult, rErr
	}
	for i := 0; i < len(listResult); i++ {
		if listResult[i].TimeStart.Before(result.TimeStart) || listResult[i].TimeEnd.After(result.TimeEnd) {
			listResult = append(listResult[:i], listResult[i+1:]...)
		}
	}
	return listResult, nil
}
*/
