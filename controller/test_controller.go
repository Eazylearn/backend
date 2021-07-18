package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
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
		Status: enum.APIStatus.Ok,
		Message: fmt.Sprintf("Test Page"),
	})
	return nil
}