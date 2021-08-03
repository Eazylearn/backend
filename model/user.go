package model

import (
	"time"
	// "encoding/json"
	// "fmt"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	// Basic information
	UserID    int64  `json:"userId" bson:"userId"`
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`

	// Relative information
	Dob     *time.Time `json:"dob,omitempty" bson:"dob,omitempty"`
	Email   string     `json:"email,omitempty" bson:"email,omitempty"`
	Address string     `json:"address,omitempty" bson:"address,omitempty"`
	Phone   string     `json:"phone,omitempty" bson:"phone,omitempty"`

	// Account and password
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Profile struct {
	// Basic information
	UserID    int64  `json:"id" bson:"id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`

	// Relative information
	Dob     *time.Time `json:"dob,omitempty" bson:"dob,omitempty"`
}

var UserDB = &db.Instance{
	CollectionName: "user",
}

func InitUserModel(db *mongo.Database) {
	UserDB.ApplyDatabase(db)
}
