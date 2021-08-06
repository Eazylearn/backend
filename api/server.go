package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
     
	"github.com/CS426FinalProject/enum"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// INITIALIZATION ////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

// Custom Server type
type APIServer struct {
	Echo *echo.Echo
}

// Custom Handler type
type Handler = func(e echo.Context) error

// Create a custom function for grouping
type ControllerFunc func(g *echo.Group) error

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

// Init a server
func InitServer() *APIServer {
	server := &APIServer{Echo: echo.New()}
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	return server
}

// Start server
func (server *APIServer) Start(port string) {
	server.Echo.Start(port)
}

// Handle methods
func (server *APIServer) SerHandler(method *enum.MethodValue, path string, h Handler) error {
	switch method.Value {
		case enum.APIMethod.GET.Value:
			server.Echo.GET(path, h)
		case enum.APIMethod.POST.Value:
			server.Echo.POST(path, h)
		case enum.APIMethod.PUT.Value:
			server.Echo.PUT(path, h)
		case enum.APIMethod.DELETE.Value:
			server.Echo.DELETE(path, h)
	}
	return nil
}

func (server *APIServer) SetGroup(group string, cf ControllerFunc) {
	g := server.Echo.Group(group)
	cf(g)
}

func GetContent(c echo.Context, template interface{}) error {
	return json.Unmarshal([]byte(GetContentText(c)), &template)
}

func Respond(context echo.Context, response *enum.APIResponse) error {
	switch response.Status {
		case enum.APIStatus.Ok:
			return context.JSON(http.StatusOK, response)
		case enum.APIStatus.Error:
			return context.JSON(http.StatusInternalServerError, response)
		case enum.APIStatus.Forbidden:
			return context.JSON(http.StatusForbidden, response)
		case enum.APIStatus.Invalid:
			return context.JSON(http.StatusBadRequest, response)
		case enum.APIStatus.NotFound:
			return context.JSON(http.StatusNotFound, response)
		case enum.APIStatus.Unauthorized:
			return context.JSON(http.StatusUnauthorized, response)
		case enum.APIStatus.Existed:
			return context.JSON(http.StatusConflict, response)
	}
	return context.JSON(http.StatusBadRequest, response)
}

func GetContentText(c echo.Context) string {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	return string(bodyBytes)
}
