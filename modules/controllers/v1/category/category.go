package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	categoryDT "gym-history/interfaces/dataTransfer/v1/category"
	"gym-history/modules/usecases/v1/category"
	"gym-history/utils"
	"net/http"
)

type CategoryController interface {
	GetAllCategory(ctx *gin.Context)
	GetCategoryById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type CategoryUseCase struct {
	useCase category.CategoryUseCase
}

func NewCategoryController(useCase category.CategoryUseCase) *CategoryUseCase {
	return &CategoryUseCase{useCase}
}

func (c *CategoryUseCase) GetAllCategory(ctx *gin.Context) {
	categories, err := c.useCase.GetAllCategory()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *CategoryUseCase) GetCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	categories, err := c.useCase.GetCategoryById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *CategoryUseCase) Create(ctx *gin.Context) {
	var input categoryDT.CategoryInput
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	categories, err := c.useCase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *CategoryUseCase) Update(ctx *gin.Context) {
	var input categoryDT.CategoryInput
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		fmt.Println("error")
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	categories, err := c.useCase.Update(id, input)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *CategoryUseCase) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	categories, err := c.useCase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}
