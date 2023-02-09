package history

import (
	"github.com/gin-gonic/gin"
	"gym-history/modules/controllers/v1/history"
)

type Route struct {
	router     *gin.Engine
	controller history.HistoryController
}

func NewHistoryRoute(controller history.HistoryController) *Route {
	return &Route{
		router:     gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	category := route.Group("/history")
	category.GET("/", r.controller.GetAllHistory)
	category.GET("/:id", r.controller.GetHistoryById)
	category.POST("/", r.controller.Create)
	category.PATCH("/:id", r.controller.Update)
	return r.router
}
