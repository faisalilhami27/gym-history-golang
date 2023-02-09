package movement

import (
	"go.mongodb.org/mongo-driver/mongo"
	movementDT "gym-history/interfaces/dataTransfer/v1/movement"
	movementModel "gym-history/modules/models/mongo/v1/movement"
	"gym-history/modules/repositories/v1/movement"
)

type repository struct {
	repository movement.MovementRepository
}

type MovementUseCase interface {
	GetAllMovement() ([]movementModel.Movement, error)
	GetMovementById(id string) (*movementModel.Movement, error)
	Create(input movementDT.MovementInput) (*mongo.InsertOneResult, error)
	Update(id string, input movementDT.MovementInput) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

func NewMovementUseCase(movementRepository movement.MovementRepository) *repository {
	return &repository{movementRepository}
}

func (r *repository) GetAllMovement() ([]movementModel.Movement, error) {
	categories, err := r.repository.GetAllMovement()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) GetMovementById(id string) (*movementModel.Movement, error) {
	data, err := r.repository.GetMovementById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *repository) Create(input movementDT.MovementInput) (*mongo.InsertOneResult, error) {
	data := movementModel.Movement{
		Name: input.Name,
	}

	result, err := r.repository.Create(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(id string, input movementDT.MovementInput) (*mongo.UpdateResult, error) {
	data := movementModel.Movement{
		Name: input.Name,
	}

	result, err := r.repository.Update(id, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	result, err := r.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
