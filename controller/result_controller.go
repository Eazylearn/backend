package controller

import (
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func ResultControllerGroup(g *echo.Group) error {
	g.GET("/submit", SubmitTestAction)
	return nil
}

//////////////////////////////////////////////////////////

// Create a test
func CreateResultAction(c echo.Context) error {

	return nil
}

func SubmitTestAction(c echo.Context) error {
	
	return nil
}
