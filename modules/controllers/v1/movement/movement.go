package movement

import (
	"fmt"
	"github.com/gin-gonic/gin"
	historyDT "gym-history/interfaces/dataTransfer/v1/movement"
	"gym-history/modules/usecases/v1/movement"
	"gym-history/utils"
	"net/http"
)

type MovementController interface {
	GetAllMovement(ctx *gin.Context)
	GetMovementById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type MovementUseCase struct {
	useCase movement.MovementUseCase
}

func NewMovementController(useCase movement.MovementUseCase) *MovementUseCase {
	return &MovementUseCase{useCase}
}

func (c *MovementUseCase) GetAllMovement(ctx *gin.Context) {
	categories, err := c.useCase.GetAllMovement()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *MovementUseCase) GetMovementById(ctx *gin.Context) {
	id := ctx.Param("id")
	categories, err := c.useCase.GetMovementById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}

func (c *MovementUseCase) Create(ctx *gin.Context) {
	var input historyDT.MovementInput
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

func (c *MovementUseCase) Update(ctx *gin.Context) {
	var input historyDT.MovementInput
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

func (c *MovementUseCase) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	categories, err := c.useCase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}
