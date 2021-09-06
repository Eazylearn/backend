package repo

import (
	"context"
	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateResult(result model.Result) error {
	//result.TotalCorrect = GetTestTotalCorrect(result.TestID, result.Answer)
	_, err := model.ResultDB.Collection.InsertOne(context.TODO(), result)
	if err != nil {
		log.Println("result_repo/CreateResult: ", err.Error())
		return err
	}

	return nil
}

func GetResultByUserID(userId int64) ([]model.Result, error) {
	list := make([]model.Result, 0)
	result, err := model.ResultDB.Collection.Find(context.TODO(), bson.M{"userId": userId})
	result.All(context.TODO(), &list)
	if err != nil {
		log.Println("result_repo/GetResultByUserID: error encoding ", err.Error())
		return list, err
	}
	return list, nil
}

func GetResultScore(result model.Result) float64 {
	var score float64
	score = 0

	totalQuestion, _ := GetTestTotalQuestion(result.TestID)
	result.TotalCorrect, _ = GetTestTotalCorrect(result.TestID, result.Answer[:])
	if totalQuestion == 0 {
		return -1
	}
	score = float64(result.TotalCorrect * 10 / totalQuestion)

	return score
}
func GetUserHistoryResult(result model.Result) ([]model.Result, error) {
	listResult := make([]model.Result, 0)
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
