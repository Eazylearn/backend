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
func TestControllerGroup(g *echo.Group) error {
	//g.GET("/", TestPage)
	g.GET("/", GetTestByIDAction)
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
func GetTestByIDAction(c echo.Context) error {
	id := c.QueryParams().Get("id")
	if id == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("test_controller/GetTestByIDAction: Empty ID"),
		})
		return nil
	}
	testID, _ := strconv.ParseInt(id, 10, 64)
	test, err := repo.GetTestByID(testID)
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
		Data:    test,
	})
	return nil
}
