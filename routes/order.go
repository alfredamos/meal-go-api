package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(server *gin.Engine) {
	//----> Apply middleware to routes
	r := server.Use(authenticate.VerifyToken)

	//----> Protected routes.
	r.GET("/orders", controllers.GetAllOrders)
	r.GET("/orders/orders-by-user-id/:userId", controllers.GetAllOrderByUserId)
	r.DELETE("/orders/delete-all-orders-by-user-id/:userId", controllers.DeleteOrderByUserId)
	r.GET("/orders/:id", controllers.GetOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)
	r.PATCH("/orders/:id/delivered", controllers.OrderDelivered)
	r.PATCH("/orders/:id/shipped", controllers.OrderShipped)
}