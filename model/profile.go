package model

import (
	"time"
)

type Profile struct {
	// Basic information
	UserID    int64  `json:"id" bson:"id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
	Address   string `json:"address" bson:"address"`
	Phone     string `json:"phone" bson:"phone"`

	// Relative information
	Dob          time.Time `json:"dob" bson:"dob"`
	AverageScore float64   `json:"avgScore" bson:"avgScore"`
	TotalTest    int       `json:"totalTest" bson:"totalTest"`

	// Account
	Username string `json:"username" bson:"username"`
}
