package model

import (
	"encoding/json"
	"time"

	// "encoding/json"
	// "fmt"

	"github.com/CS426FinalProject/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	// Basic information
	UserID    int64  `json:"userId" bson:"userId"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`

	// Relative information
	Dob     time.Time `json:"dob" bson:"dob"`
	Email   string    `json:"email" bson:"email"`
	Address string    `json:"address" bson:"address"`
	Phone   string    `json:"phone" bson:"phone"`

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
	Dob time.Time `json:"dob,omitempty" bson:"dob,omitempty"`
}

func (u User) String() string {
	ujson, _ := json.Marshal(u)
	return string(ujson)
}

var UserDB = &db.Instance{
	CollectionName: "user",
}

func InitUserModel(db *mongo.Database) {
	UserDB.ApplyDatabase(db)
}
