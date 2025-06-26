package routes

import (
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func adminRoutes(p gin.IRoutes){
	//----> Orders routes.
	p.GET("/orders", controllers.GetAllOrders)
	p.DELETE("/orders/delete-all-orders", controllers.DeleteAllOrders)
	p.PATCH("/orders/delivered/:id", controllers.OrderDelivered)
	p.PATCH("/orders/shipped/:id", controllers.OrderShipped)

	//----> Pizza routes.
	p.POST("/pizzas", controllers.CreatePizza)
	p.DELETE("/pizzas/:id", controllers.DeletePizzaById)
	p.PATCH("/pizzas/:id", controllers.EditPizzaById)
	
	//----> User routes.
	p.GET("/users", controllers.GetAllUsers)

}