package category

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gym-history/modules/models/mongo/v1/history"
)

type HistoryInput struct {
	CategoryId primitive.ObjectID  `bson:"category_id" json:"category_id"`
	MovementId primitive.ObjectID  `bson:"movement_id" json:"movement_id"`
	Date       string              `bson:"date" json:"date"`
	DetailSet  []history.DetailSet `bson:"detail_set" json:"detail_set"`
}
