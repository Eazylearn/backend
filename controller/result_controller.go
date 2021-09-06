package controller

import (
	"fmt"
	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
	"strconv"
)

// ********** Main function for managing path ********** //
func ResultControllerGroup(g *echo.Group) error {
	g.POST("/submit", SubmitTestAction)
	g.GET("/score", GetResultScoreAction)
	g.GET("/user", GetResultByUserIDAction)
	g.GET("/history", GetUserHistoryResultAction)
	return nil
}

//////////////////////////////////////////////////////////

func SubmitTestAction(c echo.Context) error {
	var body model.Result
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("result_controller.go/SubmitTestAction: Can not parse input data %s", err),
		})
	}

	iErr := repo.CreateResult(body)
	if iErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("result_controller.go/SubmitTestAction: Can not create result %s", err),
		})
	}

	return nil
}
func GetResultScoreAction(c echo.Context) error {
	var body model.Result
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("result_controller.go/GetResultScoreAction: %s", err),
		})
	}

	repo.GetResultScore(body)
	return nil

}
func GetResultByUserIDAction(c echo.Context) error {
	userId := c.QueryParams().Get("userId")
	if userId == "" {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "question_controller/GetAllQuestionByTopicIdAction: Empty ID",
		})

	}
	var id int64
	id, errConv := strconv.ParseInt(userId, 10, 0)
	if errConv != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("results_controller/GetResultByUserIDAction: %s", errConv),
		})
	}
	results, err := repo.GetResultByUserID(id)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("results_controller/GetResultByUserIDAction: %s", err),
		})
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    results,
		Message: fmt.Sprintf("Success"),
	})
	return nil
}
func GetUserHistoryResultAction(c echo.Context) error {
	var body model.Result
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("result_controller.go/GetResultScoreAction: %s", err),
		})
	}
	results, err := repo.GetUserHistoryResult(body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("results_controller/GetResultByUserIDAction: %s", err),
		})
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    results,
		Message: fmt.Sprintf("Success"),
	})
	return nil

}
