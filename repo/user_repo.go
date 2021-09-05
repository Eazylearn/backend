package repo

import (
	"context"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a user
func CreateUser(users []model.User) ([]model.User, error) {
	lastestID, _ := model.UserDB.Collection.CountDocuments(context.TODO(), bson.D{})
	list := make([]model.User, 0)
	for i := 0; i < len(users); i++ {
		lastestID = lastestID + 1
		users[i].UserID = lastestID
		_, err := model.UserDB.Collection.InsertOne(context.TODO(), users[i])
		if err != nil {
			return list, err
		}

		var user model.User
		user, findErr := GetUserByID(lastestID)
		if findErr != nil {
			return list, findErr
		}
		list = append(list, user)
	}
	return list, nil
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
		return uErr
	}
	return nil
}

func UpdateUser(users []map[string]interface{}) error {
	for i := 0; i < len(users); i++ {
		var updates []bson.M
		var filter bson.M
		for key, value := range users[i] {
			if key != "userId" {
				updates = append(updates, bson.M{"$set": bson.M{key: value}})
			} else {
				filter = bson.M{key: value}
			}
		}
		result, err := model.UserDB.Collection.UpdateOne(context.TODO(), filter, updates)
		if err != nil {
			log.Println("user_repo.go/EditUser: Find and update fail ", err.Error())
			return err
		}
		if result.MatchedCount != 0 {
			log.Println("user_repo.go/EditUser: Matched and replaced an existing document")
			return nil
		}
		if result.UpsertedCount != 0 {
			log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		}
	}
	return nil
}

func IsUserExist(username, password string) (bool, error) {
	result, err := model.UserDB.Collection.Find(context.TODO(), bson.M{"username": username, "password": password})
	list := make([]model.User, 0)
	result.All(context.TODO(), &list)
	if err != nil || len(list) == 0 {
		return false, err
	}
	return true, nil
}
