package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
)

func RootControllerGroup(g *echo.Group) error {
	g.POST("login", LoginAction)
	return nil
}

func LoginAction(c echo.Context) error {
	var body map[string]string
	err := api.GetContent(c, &body)
	if err != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "root_controller.go/LoginAction: Can not parse input data",
			Data:    false,
		})
		return err
	}
	username := body["username"]
	password := body["password"]
	exist, id, e := repo.IsUserExist(username, password)
	if e != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("root_controller.go/LoginAction: Error inserting topic %s", e.Error()),
			Data:    false,
		})
		return e
	}
	if !exist {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Unauthorized,
			Message: "Fail to login",
			Data:    exist,
		})
		return nil
	}
	// token, tErr := api.CreateToken(username, id)
	// if tErr != nil {
	// 	api.Respond(c, &enum.APIResponse{
	// 		Status:  enum.APIStatus.Error,
	// 		Message: fmt.Sprintf("root_controller.go/LoginAction: Error creating token %s", tErr.Error()),
	// 	})
	// 	return nil
	// }
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Success",
		Data:    id,
	})
	return nil
}
