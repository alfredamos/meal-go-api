package routes

import (
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func unProtectedRoutes(server *gin.Engine){
	//----> Auth-routes.
	server.POST("/auth/signup", controllers.SignupController)
	server.POST("/auth/login",controllers.LoginController)

	//----> Pizza-routes.
	server.GET("/pizzas", controllers.GetAllPizza)

}