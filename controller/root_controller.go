package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/labstack/echo/v4"
)

func RootControllerGroup(g *echo.Group) error {
	g.POST("login", LoginAction)
	return nil
}

func LoginAction(c echo.Context) error {
	var body map[string]interface{}
	err := api.GetContent(c, &body)
	fmt.Println(body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "root_controller.go/LoginAction: Can not parse input data",
		})
	}
	return nil
}
