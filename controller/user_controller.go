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
func UserControllerGroup(g *echo.Group) error {
	g.GET("/profile", GetProfileAction)
	g.GET("/find", GetUserByIDAction)
	g.POST("/create", CreateUserAction)
	g.PUT("/edit", EditUserAction)
	g.GET("/history", GetUserHistoryAction)
	return nil
}

//////////////////////////////////////////////////////////

// Create user
func CreateUserAction(c echo.Context) error {
	var body []model.User
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "user_controller.go/CreateUserAction: Can not parse input data",
		})
	}
	var users []model.User
	users, insertErr := repo.CreateUser(body)
	if insertErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("user_controller.go/CreateUserAction: Error inserting topic %s", insertErr.Error()),
		})
	}
	return api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Success",
		Data:    users,
	})
}

// Return profile
func GetProfileAction(c echo.Context) error {
	id := c.QueryParams().Get("id")
	if id == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("user_controller/GetUserByIDAction: Empty ID"),
		})
		return nil
	}
	userId, _ := strconv.ParseInt(id, 10, 64)
	profile, err := repo.GetProfileByID(userId)
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
		Data:    profile,
	})
	return nil
}

// Return a user by id
func GetUserByIDAction(c echo.Context) error {
	id := c.QueryParams().Get("id")
	if id == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("user_controller/GetUserByIDAction: Empty ID"),
		})
		return nil
	}
	userId, _ := strconv.ParseInt(id, 10, 64)
	user, err := repo.GetUserByID(userId)
	if err != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("user_controller/GetUserByIDAction: Error " + err.Error()),
		})
		return nil
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: fmt.Sprintln("Success"),
		Data:    user,
	})
	return nil
}

func UpdatePasswordAction(c echo.Context) error {
	id := c.QueryParam("id")
	pwd := c.QueryParam("pwd")
	if id == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("user_controller/UpdatePasswordAction: Empty id"),
		})
		return nil
	}
	if pwd == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("user_controller/UpdatePasswordAction: Empty password"),
		})
		return nil
	}
	userId, _ := strconv.ParseInt(id, 10, 64)
	err := repo.UpdatePassword(userId, pwd)
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
	})
	return nil
}

func EditUserAction(c echo.Context) error {
	var body []map[string]interface{}
	err := api.GetContent(c, &body)
	if err != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: "user_controller.go/EditUserAction: Can not parse input data",
		})
	}
	updateErr := repo.UpdateUser(body)
	if updateErr != nil {
		return api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("user_controller.go/EditUserAction: Error inserting topic %s", updateErr.Error()),
		})
	}
	return nil
}

func GetUserHistoryAction(c echo.Context) error {
	id := c.QueryParam("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	result, err := repo.GetResultByUserID(userId)
	if err != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("user_controller.go/GetUserHistoryAction: Error finding result %s", err.Error()),
		})
		return err
	} else {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Ok,
			Message: "Success",
			Data:    result,
		})
	}
	return nil
}
