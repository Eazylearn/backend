package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func QuestionControllerGroup(g *echo.Group) error {
	g.GET("", QuestionPage)
	return nil
}

//////////////////////////////////////////////////////////

// Testing root path of user page
func QuestionPage(c echo.Context) error {
	api.Respond(c, &enum.APIResponse{
		Status: enum.APIStatus.Ok,
		Message: fmt.Sprintf("Question Page"),
	})
	return nil
}

// Create a test
func CreateQuestionAction(c echo.Context) error {

	return nil
}