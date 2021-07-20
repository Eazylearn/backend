package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func ResultControllerGroup(g *echo.Group) error {
	g.GET("", ResultPage)
	return nil
}

//////////////////////////////////////////////////////////

// Testing root path of user page
func ResultPage(c echo.Context) error {
	api.Respond(c, &enum.APIResponse{
		Status: enum.APIStatus.Ok,
		Message: fmt.Sprintf("Result Page"),
	})
	return nil
}

// Create a test
func CreateResultAction(c echo.Context) error {

	return nil
}