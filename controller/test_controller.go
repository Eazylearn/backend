package controller

import (
	"fmt"

	"github.com/Cumbercubie/api"
	"github.com/Cumbercubie/enum"
	"github.com/Cumbercubie/model"
	"github.com/Cumbercubie/repo"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func TestControllerGroup(g *echo.Group) error {
	g.GET("/", TestPage)
	return nil
}

//////////////////////////////////////////////////////////

// Testing root path of user page
func TestPage(c echo.Context) error {
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: fmt.Sprintf("Test Page"),
	})
	return nil
}

// Create a test

func CreateTestAction(c echo.Context) error {
	var body model.Test
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "Can not parse input data",
		})
	}
	insertErr := repo.CreateTest(body)
	if insertErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("Error inserting test: %s", insertErr.Error()),
		})
	}
	return nil

}
