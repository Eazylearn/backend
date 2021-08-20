package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func TopicControllerGroup(g *echo.Group) error {
	g.POST("/create", CreateTopicAction)
	g.GET("s", GetAllTopicAction)
	// g.GET("/CreateTopics", CreateTopicsAction)
	return nil
}

//////////////////////////////////////////////////////////

// Create a topic
func CreateTopicAction(c echo.Context) error {
	var body []model.Topic
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "topic_controller.go/CreateTopicAction: Can not parse input data",
		})
	}
	insertErr := repo.CreateTopic(body)
	if insertErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("topic_controller.go/CreateTopicAction: Error inserting topic %s", insertErr.Error()),
		})
	}
	return nil
}

func GetAllTopicAction(c echo.Context) error {
	topics, err := repo.GetAllTopic()
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
		Data:    topics,
	})
	return nil
}
