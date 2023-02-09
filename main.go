package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gym-history/config"
	"gym-history/database"
	"gym-history/helpers"
	"gym-history/modules/controllers"
	"gym-history/modules/repositories"
	"gym-history/modules/usecases"
	"gym-history/routes"
	"time"
)

func init() {
	err := helpers.LoadEnv()
	if err != nil {
		return
	}
	config.InitApp()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Now().In(loc)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db := database.GetMongoClient()
	defer func(db *mongo.Client, ctx context.Context) {
		err := db.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(db, ctx)

	repositories := repositories.NewRepository(db)
	useCases := usecases.NewUseCase(repositories)
	controllers := controllers.NewController(useCases)
	router := routes.NewRoute(*controllers)
	router.Run()
}
