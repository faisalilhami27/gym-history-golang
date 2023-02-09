package category

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gym-history/modules/models/mongo/v1/category"
	"gym-history/utils"
	"time"
)

type CategoryRepository interface {
	GetAllCategory() ([]category.Category, error)
	GetCategoryById(id string) (*category.Category, error)
	Create(category category.Category) (*mongo.InsertOneResult, error)
	Update(id string, category category.Category) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type repository struct {
	db *mongo.Client
}

func NewCategoryRepository(db *mongo.Client) CategoryRepository {
	return &repository{db}
}

func (r *repository) GetAllCategory() ([]category.Category, error) {
	var result []category.Category
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("categories")
	cursor, err := collection.Find(ctx, bson.M{})

	if !utils.MongoErrorDatabaseException(err) {
		return nil, err
	}

	for cursor.Next(ctx) {
		var data category.Category
		err := cursor.Decode(&data)
		if !utils.MongoErrorDatabaseException(err) {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (r *repository) GetCategoryById(id string) (*category.Category, error) {
	var result *category.Category
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("categories")
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

func (r *repository) Create(category category.Category) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("categories")
	result, err := collection.InsertOne(ctx, category)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(id string, category category.Category) (*mongo.UpdateResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("categories")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": objectId}
	fmt.Println(category)
	data := bson.D{{"$set", bson.D{{"name", category.Name}}}}
	result, err := collection.UpdateOne(ctx, filter, data)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := r.db.Database("gym_history").Collection("categories")
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
