package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)


func RegisteredRoutes(server *gin.Engine){
	//----> Unprotected routes.
	//----> Auth-routes.
	server.POST("/auth/signup", controllers.SignupController)
	server.POST("/auth/login",controllers.LoginController)
	
	//----> Pizza-routes.
	server.GET("/pizzas", controllers.GetAllPizza)

	//----> Apply middleware for protected routes
	r := server.Use(authenticate.VerifyTokenJwt)

	//----> Protected routes.
	//----> Auth routes.
	r.PATCH("/auth/change-password", controllers.ChangePasswordController)
	r.PATCH("/auth/edit-profile", controllers.EditProfileController)
	r.POST("/auth/logout",controllers.LogoutController)
	
	//----> Cart-item routes.
	r.GET("/cart-items", controllers.GetAllCartItems)
	r.POST("/cart-items", controllers.CreateCartItem)
	r.DELETE("/cart-items/:id", controllers.DeleteCartItemById)
	r.GET("/cart-items/:id", controllers.GetCartItemById)
	r.PATCH("/cart-items/:id", controllers.EditCartItemById)

	//----> Order routes.
	r.PATCH("/orders/checkout", controllers.CheckOutOrder)
	r.GET("/orders/orders-by-user-id/:userId", controllers.GetAllOrderByUserId)
	r.DELETE("/orders/delete-all-orders-by-user-id/:userId", controllers.DeleteOrderByUserId)
	
	r.GET("/orders/:id", controllers.GetOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)
	
	//----> Pizza-routes.
	r.GET("/pizzas/:id", controllers.GetPizzaById)

	//----> User-route
	r.GET("/users/:id", controllers.GetUserById)

	//----> Admin role permitted routes middleware.
	p := server.Use(authenticate.VerifyTokenJwt, authenticate.RolePermission("Admin"))
	
	//----> Orders routes.
	p.GET("/orders", controllers.GetAllOrders)
	p.DELETE("/orders/delete-all-orders", controllers.DeleteAllOrders)
	p.PATCH("/orders/:id/delivered", controllers.OrderDelivered)
	p.PATCH("/orders/:id/shipped", controllers.OrderShipped)

	//----> Pizza routes.
	p.POST("/pizzas", controllers.CreatePizza)
	p.DELETE("/pizzas/:id", controllers.DeletePizzaById)
	p.PATCH("/pizzas/:id", controllers.EditPizzaById)
	
	//----> User routes.
	p.GET("/users", controllers.GetAllUsers)
	p.DELETE("/users/:id", controllers.DeleteUserById)
}