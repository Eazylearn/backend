package controller

import (
	"fmt"
	"strconv"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func SubjectControllerGroup(g *echo.Group) error {
	g.POST("/create", CreateSubjectAction)
	g.GET("s", GetAllSubjectAction)
	g.GET("", GetSubjectByIDAction)
	return nil
}

//////////////////////////////////////////////////////////

// Create a topic
func CreateSubjectAction(c echo.Context) error {
	var body model.Subject
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "subject_controller.go/CreateSubjectAction: Can not parse input data",
		})
	}
	insertErr := repo.CreateSubject(body)
	if insertErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("subject_controller.go/CreateSubjectAction: Error inserting subject %s", insertErr.Error()),
		})
	}
	return nil
}

func GetSubjectByIDAction(c echo.Context) error {
	id := c.QueryParams().Get("id")
	if id == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("subject_controller/GetSubjectByIDAction: Empty ID"),
		})
		return nil
	}
	subjectId, _ := strconv.ParseInt(id, 10, 64)
	subject, err := repo.GetSubjectByID(subjectId)
	if err != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf(err.Error()),
		})
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: fmt.Sprintln("Success"),
		Data:    subject,
	})
	return nil
}

func GetAllSubjectAction(c echo.Context) error {
	subjects, err := repo.GetAllSubject()
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
		Data:    subjects,
	})
	return nil
}
