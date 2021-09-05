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
func TestControllerGroup(g *echo.Group) error {
	g.POST("/create", CreateTestAction)
	g.POST("/createbyBE", CreateTestByBEAction)
	g.GET("/", GetAllTestByQueryAtcion)

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
func CreateTestByBEAction(c echo.Context) error {
	var body model.PostTest
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "Can not parse input data",
		})
	}
	insertErr := repo.CreateTestByBE(body)
	if insertErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("Error inserting test: %s", insertErr.Error()),
		})
	}
	return nil
}
func GetAllTestByQueryAtcion(c echo.Context) error {
	var input model.Test
	param := c.QueryParams().Get("q")
	if param == "" {
		param = "{}"
	}
	paramErr := json.Unmarshal([]byte(param), &input)
	if paramErr != nil {
		return nil
	}
	test, err := repo.GetAllTestByQuery(&input)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("test_controller/GetAllTestByQueryAtcion: %s", err),
		})
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    test,
		Message: fmt.Sprintf("Success"),
	})
	return nil
}
