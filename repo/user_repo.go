package repo

import (
	"context"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a user
func CreateUser(user []model.User) error {
	// return &User{
	//		userId: 1,
	//		firstName: "first name",
	//		lastName: "last name",
	//		dob: yyyy-mm-dd,
	//		email: "abc@gmail.com"
	//		phone: "0123456789",
	//		username: "abc",
	//		password: "123"
	//	}
	for i := 0; i < len(user); i++ {
		_, err := model.UserDB.Collection.InsertOne(context.TODO(), user[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// Return profile by id
func GetProfileByID(id int64) (model.Profile, error) {
	var profile model.Profile
	var user model.User
	err := model.UserDB.Collection.FindOne(context.TODO(), bson.D{{"userId", id}}).Decode(&user)
	if err != nil {
		log.Println("user_repo.go/FindUserByID: Error finding", err.Error())
		return profile, err
	}
	profile.UserID = user.UserID
	profile.FirstName = user.FirstName
	profile.LastName = user.LastName
	profile.Dob = user.Dob
	return profile, nil
}

// Return user by id
func GetUserByID(id int64) (model.User, error) {
	var user model.User
	err := model.UserDB.Collection.FindOne(context.TODO(), bson.D{{"userId", id}}).Decode(&user)
	if err != nil {
		log.Println("user_repo.go/FindUserByID: Error finding", err.Error())
		return user, err
	}
	return user, nil
}

// Return all users by first name
func GetUserByFirstname(firstName string) ([]model.User, error) {
	list := make([]model.User, 0)
	result := model.UserDB.Collection.FindOne(context.TODO(), bson.M{"firstName": firstName})
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		log.Println("user_repo/GetUserByFirstname: ", err.Error())
		return list, err
	}
	list = append(list, user)
	return list, nil
}

// Return username
func GetUsername(id int) (string, error) {
	result := model.UserDB.Collection.FindOne(context.TODO(), bson.M{"userId": id})
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		log.Println("user_repo/GetUsername: ", err.Error())
		return "", err
	}
	return user.Username, nil
}

// Return password
func GetPassword(id int) (string, error) {
	result := model.UserDB.Collection.FindOne(context.TODO(), bson.M{"userId": id})
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		log.Println("user_repo/GetPassword: ", err.Error())
		return "", err
	}
	return user.Password, nil
}

// Delete a user
func DeleteUserByID(id int) error {

	return nil
}

// Update password
func UpdatePassword(id int64, pwd string) error {
	filter := bson.D{{"userId", id}}
	update := bson.D{{"$set", bson.D{{"password", pwd}}}}
	var updateDoc bson.M
	uErr := model.UserDB.Collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updateDoc)
	if uErr != nil {
		log.Println("user_repo.go/UpdatePassword: Update fail", uErr.Error())
	}
	return nil
}
