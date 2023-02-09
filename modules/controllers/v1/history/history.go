package history

import (
	"fmt"
	"github.com/gin-gonic/gin"
	historyDT "gym-history/interfaces/dataTransfer/v1/history"
	"gym-history/modules/usecases/v1/history"
	"gym-history/utils"
	"net/http"
)

type HistoryController interface {
	GetAllHistory(ctx *gin.Context)
	GetHistoryById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type HistoryUseCase struct {
	useCase history.HistoryUseCase
}

func NewHistoryController(useCase history.HistoryUseCase) *HistoryUseCase {
	return &HistoryUseCase{useCase}
}

func (c *HistoryUseCase) GetAllHistory(ctx *gin.Context) {
	categories, err := c.useCase.GetAllHistory()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *HistoryUseCase) GetHistoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	categories, err := c.useCase.GetHistoryById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *HistoryUseCase) Create(ctx *gin.Context) {
	var input historyDT.HistoryInput
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

func (c *HistoryUseCase) Update(ctx *gin.Context) {
	var input historyDT.HistoryInput
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
