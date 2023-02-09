package history

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gym-history/modules/models/mongo/v1/category"
	"gym-history/modules/models/mongo/v1/movement"
	"time"
)

type Date string

func (d *Date) UnmarshalJSON(bytes []byte) error {
	dd, err := time.Parse(`"2006-01-02T15:04:05.000+0000"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(dd.Format("01/02/2006"))

	return nil
}

type History struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CategoryId primitive.ObjectID `bson:"category_id,omitempty" json:"category_id,omitempty"`
	MovementId primitive.ObjectID `bson:"movement_id" json:"movement_id"`
	Date       string             `bson:"date" json:"date"`
	DetailSet  []DetailSet        `bson:"detail_set" json:"detail_set"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type HistoryFormat struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Category   category.Category  `bson:"category,omitempty" json:"category"`
	MovementId movement.Movement  `bson:"movement,omitempty" json:"movement"`
	Date       string             `bson:"date" json:"date"`
	DetailSet  []DetailSet        `bson:"detail_set" json:"detail_set"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type DetailSet struct {
	Set         int    `bson:"set" json:"set"`
	Repetition  int    `bson:"rep" json:"repetition"`
	Weight      int    `bson:"weight" json:"weight"`
	Description string `bson:"description" json:"description"`
}
