package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gym-history/modules/controllers"
	"gym-history/routes/v1/category"
	"gym-history/routes/v1/history"
	"gym-history/routes/v1/movement"
	"os"
)

type Route struct {
	controller controllers.Controller
}

func NewRoute(controller controllers.Controller) *Route {
	return &Route{
		controller: controller,
	}
}

func (r *Route) Run() {
	router := gin.New()
	router.Use(cors.Default())
	router.Static("assets", "./public")
	group := router.Group("/api/v1")
	categoryRoute := category.NewCategoryRoute(r.controller.CategoryController)
	categoryRoute.Start(group)

	movementRoute := movement.NewMovementRoute(r.controller.MovementController)
	movementRoute.Start(group)

	historyRoute := history.NewHistoryRoute(r.controller.HistoryController)
	historyRoute.Start(group)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := router.Run(port)
	if err != nil {
		return
	}
}
