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
	UserID    int64  `json:"userId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`

	// Relative information
	Dob     *time.Time `json:"dob,omitempty"`
	Email   string     `json:"email,omitempty"`
	Address string     `json:"address,omitempty"`
	Phone   string     `json:"phone,omitempty"`

	// Account and password
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
}

// func (u User) String() string {
// 	ujson, _ := json.Marshal(u)
// 	return fmt.Sprintf(ujson)
// }

var UserDB = &db.Instance{
	CollectionName: "user",
}

func InitUserModel(db *mongo.Database) {
	UserDB.ApplyDatabase(db)
}
