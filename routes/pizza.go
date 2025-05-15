package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPizzaRoutes(server *gin.Engine) {
	//----> Unprotected routes.
	server.GET("/pizzas", controllers.GetAllPizza)

	//----> Protected routes
	//----> Apply middleware to routes
	r := server.Use(authenticate.VerifyToken)
	
	r.POST("/pizzas", controllers.CreatePizza)
	r.GET("/pizzas/:id", controllers.GetPizzaById)
	r.DELETE("/pizzas/:id", controllers.DeletePizzaById)
	r.PATCH("/pizzas/:id", controllers.EditPizzaById)
}