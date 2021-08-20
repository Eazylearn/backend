package repo

import (
	"context"
	"log"

	"github.com/CS426FinalProject/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTopic(topic []model.Topic) error {
	// return &Topic{TopicId: 1, Name: "New topic"}
	for i := 0; i < len(topic); i++ {
		_, err := model.TopicDB.Collection.InsertOne(context.TODO(), topic[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllTopic() ([]model.Topic, error) {
	list := make([]model.Topic, 0)
	result, err := model.TopicDB.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("topic_repo/GetAllTopic: ", err.Error())
		return list, err
	}
	result.All(context.TODO(), &list)
	return list, nil
}
