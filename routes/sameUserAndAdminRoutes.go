package routes

import (
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func sameUserAndAdminRoutes(r gin.IRoutes) {
	//----> Order routes.
	r.GET("/orders/orders-by-user-id/:userId", controllers.GetAllOrderByUserId)
	r.DELETE("/orders/delete-all-orders-by-user-id/:userId", controllers.DeleteOrderByUserId)
	
	//----> User-route
	r.GET("/users/:id", controllers.GetUserById)
	r.DELETE("/users/:id", controllers.DeleteUserById)
}