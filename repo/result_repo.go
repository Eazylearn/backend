package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateResult(result model.Result) error {
	// return &Result{
	//		resultId: 1,
	//		timeStart: ,
	//		timeEnd: ,
	//		totalCorrect: 0,
	//		totalIncorrect: 0,
	//		answeredQuestion: 0,
	//		totalTime: ,
	//		userId: 1,
	//		testId: 1
	//				}
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
