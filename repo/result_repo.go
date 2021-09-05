package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateResult(result model.Result) error {
	_, err := model.ResultDB.Collection.InsertOne(context.TODO(), result)
	if err != nil {
		return err
	}
	return nil
}

func GetResultByUserID(userId int64) ([]model.Result, error) {
	list := make([]model.Result, 0)
	result, err := model.ResultDB.Collection.Find(context.TODO(), bson.M{"userId": userId})
	result.All(context.TODO(), &list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func GetResultScore(result model.Result) float64 {
	var score float64
	score = 0

	totalQuestion := GetTestTotalQuestion(result.TestID)

	score = float64(result.TotalCorrect) / float64(totalQuestion)

	return score
}
