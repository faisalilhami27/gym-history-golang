package movement

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gym-history/modules/models/mongo/v1/movement"
	"gym-history/utils"
	"time"
)

type MovementRepository interface {
	GetAllMovement() ([]movement.Movement, error)
	GetMovementById(id string) (*movement.Movement, error)
	Create(movement movement.Movement) (*mongo.InsertOneResult, error)
	Update(id string, movement movement.Movement) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type repository struct {
	db *mongo.Client
}

func NewMovementRepository(db *mongo.Client) MovementRepository {
	return &repository{db}
}

func (r *repository) GetAllMovement() ([]movement.Movement, error) {
	var result []movement.Movement
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("movements")
	cursor, err := collection.Find(ctx, nil)

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	for cursor.Next(ctx) {
		var data movement.Movement
		err := cursor.Decode(&data)
		if !utils.MongoErrorDatabaseException(err) {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (r *repository) GetMovementById(id string) (*movement.Movement, error) {
	var result *movement.Movement
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("movements")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Create(movement movement.Movement) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("movements")
	result, err := collection.InsertOne(ctx, movement)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(id string, movement movement.Movement) (*mongo.UpdateResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("movements")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	data := bson.D{{"$set", bson.D{{"name", movement.Name}}}}
	result, err := collection.UpdateOne(ctx, filter, data)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("movements")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}
