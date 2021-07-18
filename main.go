package main

import (
	"fmt"
	"os"

	"github.com/CS426FinalProject/api"
	"github.com/CS426FinalProject/controller"
	"github.com/CS426FinalProject/db"
	"github.com/CS426FinalProject/enum"
	"github.com/CS426FinalProject/model"
	"github.com/CS426FinalProject/repo"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	server := api.InitServer()
	godotenv.Load(".env")
	DB_URI := os.Getenv("DB_URI")
	server.SetGroup("/", hello)
}