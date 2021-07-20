package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
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