package main

import (
	// "os"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/controller"
	// "github.com/CS426FinalProject/db"
	// "github.com/CS426FinalProject/enum"
	// "github.com/CS426FinalProject/model"
	// "github.com/CS426FinalProject/repo"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main() {
	server := api.InitServer()
	godotenv.Load(".env")
	// DB_URI := os.Getenv("DB_URI")
	server.SetGroup("/", controller.RootControllerGroup)
	server.SetGroup("/user", controller.UserControllerGroup)
	
	server.Start(":8081")
}