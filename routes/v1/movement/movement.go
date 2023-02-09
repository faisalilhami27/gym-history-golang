package movement

import (
	"github.com/gin-gonic/gin"
	"gym-history/modules/controllers/v1/movement"
)

type Route struct {
	router     *gin.Engine
	controller movement.MovementController
}

func NewMovementRoute(controller movement.MovementController) *Route {
	return &Route{
		router:     gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	category := route.Group("/movement")
	category.GET("/", r.controller.GetAllMovement)
	category.GET("/:id", r.controller.GetMovementById)
	category.POST("/", r.controller.Create)
	category.PATCH("/:id", r.controller.Update)
	category.DELETE("/:id", r.controller.Delete)
	return r.router
}
