package controller

import (
	"fmt"
	"strconv"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/repo"
	"github.com/labstack/echo/v4"
)

// ********** Main function for managing path ********** //
func UserControllerGroup(g *echo.Group) error {
	g.GET("/profile", GetProfileAction)
	g.GET("/user", GetUserByIDAction)
	//g.GET("/user", GetUserByFirstnameAction)
	return nil
}

//////////////////////////////////////////////////////////

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
			Message: fmt.Sprintf(err.Error()),
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

// Return all users by first name
func GetUserByFirstnameAction(c echo.Context) error {
	firstName := c.QueryParams().Get("firstName")
	if firstName == "" {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintln("user_controller/GetUserByFirstnameAction: Empty first name"),
		})
		return nil
	}
	user, err := repo.GetUserByFirstname(firstName)
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
