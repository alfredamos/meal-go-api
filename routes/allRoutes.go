package routes

import "github.com/gin-gonic/gin"

func RegisterAllRoutes(server *gin.Engine) {
	RegisterAuthRoutes(server)
	RegisterOrderRoutes(server)
	RegisterPizzaRoutes(server)
	RegisterUserRoutes(server)
}
