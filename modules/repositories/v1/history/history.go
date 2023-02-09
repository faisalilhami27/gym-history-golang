package history

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gym-history/modules/models/mongo/v1/history"
	"gym-history/utils"
	"time"
)

type HistoryRepository interface {
	GetAllHistory() ([]history.HistoryFormat, error)
	GetHistoryById(id string) (*history.History, error)
	Create(history history.History) (*mongo.InsertOneResult, error)
	Update(id string, history history.History) (*mongo.UpdateResult, error)
}

type repository struct {
	db *mongo.Client
}

func NewHistoryRepository(db *mongo.Client) HistoryRepository {
	return &repository{db}
}

func (r *repository) GetAllHistory() ([]history.HistoryFormat, error) {
	var result []history.HistoryFormat
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("histories")
	category := bson.D{{"$lookup", bson.D{{"from", "categories"}, {"localField", "category_id"}, {"foreignField", "_id"}, {"as", "category"}}}}
	movement := bson.D{{"$lookup", bson.D{{"from", "movements"}, {"localField", "movement_id"}, {"foreignField", "_id"}, {"as", "movement"}}}}
	unwindCategory := bson.D{{"$unwind", bson.D{{"path", "$category"}, {"preserveNullAndEmptyArrays", false}}}}
	unwindMovement := bson.D{{"$unwind", bson.D{{"path", "$movement"}, {"preserveNullAndEmptyArrays", false}}}}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{category, movement, unwindMovement, unwindCategory})

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		panic(err)
	}

	return result, nil
}

func (r *repository) GetHistoryById(id string) (*history.History, error) {
	var result *history.History
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("histories")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	return result, nil
}

func (r *repository) Create(history history.History) (*mongo.InsertOneResult, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("histories")
	history.CreatedAt = time.Now().In(location)
	history.UpdatedAt = time.Now().In(location)
	result, err := collection.InsertOne(ctx, history)

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(id string, history history.History) (*mongo.UpdateResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("histories")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": history}
	result, err := collection.UpdateOne(ctx, filter, update)

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	return result, nil
}
