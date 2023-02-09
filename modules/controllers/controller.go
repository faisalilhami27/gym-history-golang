package controllers

import (
	"gym-history/modules/controllers/v1/category"
	"gym-history/modules/controllers/v1/history"
	"gym-history/modules/controllers/v1/movement"
	"gym-history/modules/usecases"
)

type Controller struct {
	CategoryController category.CategoryController
	MovementController movement.MovementController
	HistoryController  history.HistoryController
}

func NewController(useCase *usecases.UseCase) *Controller {
	categoryController := category.NewCategoryController(useCase.CategoryUseCase)
	movementController := movement.NewMovementController(useCase.MovementUseCase)
	historyController := history.NewHistoryController(useCase.HistoryUseCase)
	return &Controller{
		CategoryController: categoryController,
		MovementController: movementController,
		HistoryController:  historyController,
	}
}
