package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	//----> Apply middleware to routes
	list := make([]string, 0)
	list = append(list, "Admin")
	//list = append(list, "Customer")
	r := server.Use(authenticate.VerifyToken, authenticate.RolePermission(list))

	//----> Protected routes.
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetPizzaById)
	r.DELETE("/users/:id", controllers.DeleteUserById)
}