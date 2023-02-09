package history

import (
	"go.mongodb.org/mongo-driver/mongo"
	historyDT "gym-history/interfaces/dataTransfer/v1/history"
	historyModel "gym-history/modules/models/mongo/v1/history"
	"gym-history/modules/repositories/v1/history"
)

type repository struct {
	repository history.HistoryRepository
}

type HistoryUseCase interface {
	GetAllHistory() ([]historyModel.HistoryFormat, error)
	GetHistoryById(id string) (*historyModel.History, error)
	Create(input historyDT.HistoryInput) (*mongo.InsertOneResult, error)
	Update(id string, input historyDT.HistoryInput) (*mongo.UpdateResult, error)
}

func NewHistoryUseCase(historyRepository history.HistoryRepository) *repository {
	return &repository{historyRepository}
}

func (r *repository) GetAllHistory() ([]historyModel.HistoryFormat, error) {
	categories, err := r.repository.GetAllHistory()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) GetHistoryById(id string) (*historyModel.History, error) {
	data, err := r.repository.GetHistoryById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *repository) Create(input historyDT.HistoryInput) (*mongo.InsertOneResult, error) {
	data := historyModel.History{
		CategoryId: input.CategoryId,
		MovementId: input.MovementId,
		Date:       input.Date,
		DetailSet:  input.DetailSet,
	}

	result, err := r.repository.Create(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(id string, input historyDT.HistoryInput) (*mongo.UpdateResult, error) {
	data := historyModel.History{
		CategoryId: input.CategoryId,
		MovementId: input.MovementId,
		Date:       input.Date,
		DetailSet:  input.DetailSet,
	}

	result, err := r.repository.Update(id, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
