package repo

import (
	"context"

	"github.com/CS426FinalProject/model"
)

func CreateTopic(topic model.Topic) error {
	// return &Topic{TopicId: 1, Name: "New topic"}
	_, err := model.TopicDB.Collection.InsertOne(context.TODO(), topic)
	if err != nil {
		return err
	}
	return nil
}