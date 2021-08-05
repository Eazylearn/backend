package repo

import (
	"context"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a user
func CreateUser(user model.User) (model.User, error) {
	// return &User{
	//		userId: 1,
	//		firstName: "first name",
	//		lastName: "last name",
	//		dob: yyyy-mm-dd,
	//		email: "abc@gmail.com"
	//		phone: "0123456789",
	//		account: "abc",
	//		password: "123"
	//	}
	var newUser model.User
	_, err := model.UserDB.Collection.InsertOne(context.TODO(), user)
	// Query new user
	// ...
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// Return profile by id
func GetProfileByID(id int64) (model.Profile, error) {
	var user model.User
	var profile model.Profile
	result, qErr := model.UserDB.Collection.Find(context.TODO(), bson.M{"userId": id})
	if qErr != nil {
		log.Println("user_repo.go/FindUserByID: Error finding", qErr.Error())
		return profile, qErr
	}
	profile.UserID = user.UserID
	profile.FirstName = user.FirstName
	profile.LastName = user.LastName
	profile.Dob = user.Dob
	err := result.All(context.TODO(), &profile)
	if err != nil {
		log.Println("user_repo.go/FindUserByID: Error encoding", err.Error())
		return profile, err
	}
	return profile, nil
}

// Return user by id
func GetUserByID(id int64) (model.User, error) {
	var user model.User
	result, qErr := model.UserDB.Collection.Find(context.TODO(), bson.M{"userId": id})
	if qErr != nil {
		log.Println("user_repo.go/FindUserByID: Error finding", qErr.Error())
		return user, qErr
	}
	err := result.All(context.TODO(), &user)
	if err != nil {
		log.Println("user_repo.go/FindUserByID: Error encoding", err.Error())
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

