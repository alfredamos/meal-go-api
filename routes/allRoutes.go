package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/gin-gonic/gin"
)


func RegisteredRoutes(server *gin.Engine){
	//----> Unprotected routes.
	unProtectedRoutes(server)
	
	//----> Apply middleware for protected routes
	r := server.Group("/api").Use(authenticate.VerifyTokenJwt)

	//----> Protected routes.
	protectedRoutes(r)
	
	//----> Admin role permitted routes middleware.
	p := server.Group("/api").Use(authenticate.VerifyTokenJwt, authenticate.RolePermission("Admin"))
	
	//----> Admin routes
	adminRoutes(p)
}