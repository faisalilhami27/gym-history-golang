package usecases

import (
	"gym-history/modules/repositories"
	"gym-history/modules/usecases/v1/category"
	"gym-history/modules/usecases/v1/history"
	"gym-history/modules/usecases/v1/movement"
)

type UseCase struct {
	CategoryUseCase category.CategoryUseCase
	MovementUseCase movement.MovementUseCase
	HistoryUseCase  history.HistoryUseCase
}

func NewUseCase(repository *repositories.Repository) *UseCase {
	categoryUseCase := category.NewCategoryUseCase(repository.CategoryRepository)
	historyUseCase := history.NewHistoryUseCase(repository.HistoryRepository)
	movementUseCase := movement.NewMovementUseCase(repository.MovementRepository)
	return &UseCase{
		CategoryUseCase: categoryUseCase,
		MovementUseCase: movementUseCase,
		HistoryUseCase:  historyUseCase,
	}
}
