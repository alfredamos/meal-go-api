package routes

import (
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func ownerAndAdminRoutes(r gin.IRoutes) {
	//----> Order routes.
	r.GET("/orders/:id", controllers.GetOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)
	
}