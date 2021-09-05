package controller

import (
	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func ResultControllerGroup(g *echo.Group) error {
	g.GET("/submit", SubmitTestAction)
	return nil
}

//////////////////////////////////////////////////////////

func SubmitTestAction(c echo.Context) error {
	var body model.Result
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "result_controller.go/SubmitTestAction: Can not parse input data",
		})
	}

	iErr := repo.CreateResult(body)
	if iErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "result_controller.go/SubmitTestAction: Can not create result",
		})
	}

	return nil
}
