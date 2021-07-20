package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a user
func CreateUser(user model.User) error {
	// return &User{
	//		userId: 1,
	//		firstName: "first name",
	//		lastName: "last name",
	//		dob: ,
	//		email: "abc@gmail.com"
	//		phone: "0123456789",
	//		account: "abc",
	//		password: "123"
	//	}
	_, err := model.UserDB.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// Return user by id
func FindUserByID(id int) (model.User, error) {
	var user model.User
	result, qErr := model.UserDB.Collection.Find(context.TODO(), bson.M{})
	if qErr != nil {
		log.Println("user_repo.go/FindUserByID: Error finding", qErr.Error())
		return user, qErr
	}
	fmt.Println(result == nil)
	err := result.All(context.TODO(), &user)
	if err != nil {
		log.Println("user_repo.go/FindUserByID: Error encoding", err.Error())
		return user, err
	}
	return user, nil
}