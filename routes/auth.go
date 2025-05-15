package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)


func RegisterAuthRoutes(server *gin.Engine){
	//----> Unprotected routes.
	server.GET("/", controllers.Home)
	server.POST("/auth/signup", controllers.SignupController)
	server.POST("/auth/login",controllers.LoginController)
	
	//----> Apply middleware to routes
	r := server.Use(authenticate.VerifyToken)
	
	//----> Protected routes.
	r.PATCH("/auth/change-password", controllers.ChangePasswordController)
	r.PATCH("/auth/edit-profile", controllers.EditProfileController)
	r.POST("/auth/logout",controllers.LogoutController)

	//----> Roles permitted routes
}
