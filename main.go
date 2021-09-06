package main

import (
	"os"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/controller"
	"github.com/CS426FinalProject/db"

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
	// setMiddleware(server)
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

// Connect to database
func onDBConnected(c *mongo.Database) {
	model.InitUserModel(c)
	model.InitTestDB(c)
	model.InitResultDB(c)
	model.InitQuestionDB(c)
	model.InitTopicDB(c)
	model.InitSubjectDB(c)
}
