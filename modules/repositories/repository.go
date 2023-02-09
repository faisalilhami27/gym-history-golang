package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gym-history/modules/repositories/v1/category"
	"gym-history/modules/repositories/v1/history"
	"gym-history/modules/repositories/v1/movement"
)

type Repository struct {
	CategoryRepository category.CategoryRepository
	MovementRepository movement.MovementRepository
	HistoryRepository  history.HistoryRepository
}

func NewRepository(db *mongo.Client) *Repository {
	categoryRepository := category.NewCategoryRepository(db)
	movementRepository := movement.NewMovementRepository(db)
	historyRepository := history.NewHistoryRepository(db)
	return &Repository{
		CategoryRepository: categoryRepository,
		MovementRepository: movementRepository,
		HistoryRepository:  historyRepository,
	}
}
