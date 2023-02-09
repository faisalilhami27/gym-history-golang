package category

import (
	"go.mongodb.org/mongo-driver/mongo"
	categoryDT "gym-history/interfaces/dataTransfer/v1/category"
	categoryModel "gym-history/modules/models/mongo/v1/category"
	"gym-history/modules/repositories/v1/category"
)

type repository struct {
	repository category.CategoryRepository
}

type CategoryUseCase interface {
	GetAllCategory() ([]categoryModel.Category, error)
	GetCategoryById(id string) (*categoryModel.Category, error)
	Create(input categoryDT.CategoryInput) (*mongo.InsertOneResult, error)
	Update(id string, input categoryDT.CategoryInput) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

func NewCategoryUseCase(categoryRepository category.CategoryRepository) *repository {
	return &repository{categoryRepository}
}

func (r *repository) GetAllCategory() ([]categoryModel.Category, error) {
	categories, err := r.repository.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) GetCategoryById(id string) (*categoryModel.Category, error) {
	data, err := r.repository.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *repository) Create(input categoryDT.CategoryInput) (*mongo.InsertOneResult, error) {
	data := categoryModel.Category{
		Name: input.Name,
	}

	result, err := r.repository.Create(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(id string, input categoryDT.CategoryInput) (*mongo.UpdateResult, error) {
	data := categoryModel.Category{
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
