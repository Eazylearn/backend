package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/controller"
	"github.com/CS426FinalProject/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	// "github.com/CS426FinalProject/repo"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	server := api.InitServer()
	godotenv.Load(".env")
	DB_URI := os.Getenv("DB_URI")
	PORT := ":" + os.Getenv("PORT")
	AppDB := db.CreateUniversalDB(DB_URI, "eazylearn")
	onDBConnected(AppDB)
	createPath(server)
	setMiddleware(server)
	server.Start(PORT)
}

// Create path
func createPath(server *api.APIServer) {
	server.SetGroup("/", controller.RootControllerGroup)
	server.SetGroup("/user", controller.UserControllerGroup)
	server.SetGroup("/test", controller.TestControllerGroup)
	server.SetGroup("/topic", controller.TopicControllerGroup)
	server.SetGroup("/question", controller.QuestionControllerGroup)
	server.SetGroup("/result", controller.ResultControllerGroup)
	server.SetGroup("/subject", controller.SubjectControllerGroup)
}

// Set middleware
func setMiddleware(server *api.APIServer) {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

	r := server.Echo.Group("/user")

	config := middleware.JWTConfig{
		TokenLookup: api.GetHeaderText(server.Echo.AcquireContext()),
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				signingKey := []byte("secret")
				return signingKey, nil
			}

			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	r.Use(middleware.JWTWithConfig(config))
}

// Connect to database
func onDBConnected(c *mongo.Database) {
	model.InitUserModel(c)
	model.InitTestDB(c)
	model.InitResultDB(c)
	model.InitQuestionDB(c)
	model.InitTopicDB(c)
	model.InitSubjectDB(c)
}
