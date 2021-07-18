package controller

import (
	"fmt"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/enum"
	"github.com/labstack/echo/v4"
)

func RootControllerGroup(g *echo.Group) error {
	g.GET("", hello)
	g.POST("", hello)
	g.HEAD("", hello)
	return nil
}

func hello(c echo.Context) error {
	api.Respond(c, &enum.APIResponse{
		Status: enum.APIStatus.Ok,
		Message: fmt.Sprintf("Hello World!"),
	})
	return nil
}