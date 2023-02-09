package category

import (
	"github.com/gin-gonic/gin"
	"gym-history/modules/controllers/v1/category"
)

type Route struct {
	router     *gin.Engine
	controller category.CategoryController
}

func NewCategoryRoute(controller category.CategoryController) *Route {
	return &Route{
		router:     gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	category := route.Group("/category")
	category.GET("/", r.controller.GetAllCategory)
	category.GET("/:id", r.controller.GetCategoryById)
	category.POST("/", r.controller.Create)
	category.PATCH("/:id", r.controller.Update)
	category.DELETE("/:id", r.controller.Delete)
	return r.router
}
