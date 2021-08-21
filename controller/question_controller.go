package controller

import (
	"encoding/json"
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func QuestionControllerGroup(g *echo.Group) error {
	g.GET("/", GetAllQuestionAction)
	//g.GET("/GetAllQuestionByTopicId", GetAllQuestionByTopicIdAction)
	//g.GET("/GetQuestioByIndex", GetQuestioByIndexAction)
	return nil
}

func GetAllQuestionAction(c echo.Context) error {
	//question, err := repo.GetAllQuestion()
	var input model.GetQuestionRequest
	param := c.QueryParams().Get("q")
	if param == "" {
		param = "{}"
	}
	paramErr := json.Unmarshal([]byte(param), &input)
	if paramErr != nil {
		return nil
	}
	question, err := repo.GetAllQuestionByQuery(&input)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("question_controller/GetAllQuestionAction: %s", err),
		})
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    question,
		Message: fmt.Sprintf("Success"),
	})
	return nil
}

func GetAllQuestionByTopicIdAction(c echo.Context) error {
	topicId := c.QueryParams().Get("topicId")
	if topicId == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("question_controller/GetAllQuestionByTopicIdAction: Empty ID"),
		})
		return nil
	}
	question, err := repo.GetAllQuestionByTopicId(topicId)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("Error: %s", err),
		})
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    question,
		Message: fmt.Sprintf("Success"),
	})
	return nil
}
func GetQuestioByIndexAction(c echo.Context) error {
	index := c.QueryParams().Get("index")
	if index == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("question_controller/GetQuestioByIndexAction: Empty index"),
		})
		return nil
	}
	question, err := repo.GetQuestioByIndex(index)
	if err != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf(err.Error()),
		})
		return nil
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: fmt.Sprintln("Success"),
		Data:    question,
	})
	return nil
}

//func GetAllQuestionByQuerryAction()
